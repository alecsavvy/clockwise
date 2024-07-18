package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alecsavvy/clockwise/core"
	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/graph"
	"github.com/alecsavvy/clockwise/utils"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed templates/*
var embeddedFiles embed.FS

func run() error {
	// app level context
	ctx := context.Background()

	// logger setup
	logger := utils.NewLogger(nil)
	logger.Info("good morning")

	// db setup
	pgConnectionString := os.Getenv("pgConnectionString")

	err := db.RunMigrations(logger, pgConnectionString)
	if err != nil {
		return utils.AppError("could not complete database migrations", err)
	}

	pool, err := pgxpool.New(ctx, pgConnectionString)
	if err != nil {
		return utils.AppError("failure to create db pool", err)
	}
	defer pool.Close()

	// config setup
	homeDir := "./cmt-home"

	// only keep last 100 blocks after successful indexing
	// move this to a .env config or node type "archive", "service", or something
	coreApp := core.NewCore(logger, pool, 100)
	node, err := core.NewNode(logger, homeDir, pgConnectionString, coreApp)
	if err != nil {
		return utils.AppError("failure to init chain", err)
	}

	// graphql setup
	gqlResolver := graph.NewResolver(logger, coreApp, pool)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: gqlResolver}))
	queryHandler := func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	}

	// Add WebSocket support for subscriptions
	srv.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		KeepAlivePingInterval: 10 * time.Second,
	})

	// HTTP transport for queries and mutations
	srv.AddTransport(transport.POST{})

	gqlPlayground := playground.Handler("GraphQL playground", "/query")
	graphiqlHandler := func(c echo.Context) error {
		gqlPlayground.ServeHTTP(c.Response(), c.Request())
		return nil
	}

	htmlFS, err := fs.Sub(embeddedFiles, "templates")
	if err != nil {
		return err
	}
	htmlTemplates := echo.WrapHandler(http.FileServerFS(htmlFS))

	// web server setup
	e := echo.New()

	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/graphiql", graphiqlHandler)
	e.Any("/query", queryHandler)
	e.GET("/", htmlTemplates)

	// run all the processes
	var wg sync.WaitGroup

	wg.Add(3)

	// run chain
	go func() {
		defer wg.Done()
		node.Start()

		defer func() {
			node.Stop()
			node.Wait()
		}()

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
	}()

	// run web server
	go func() {
		defer wg.Done()
		err = e.Start("0.0.0.0:26659")
		if err != nil {
			logger.Error("web server crashed", err)
		}
	}()

	// // run pubsub listener
	go func() {
		defer wg.Done()
		if err := coreApp.Run(node); err != nil {
			logger.Error("core app crashed", err)
		}
	}()

	wg.Wait()
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println("fatal error: ", err)
	}

}
