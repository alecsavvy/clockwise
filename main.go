package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alecsavvy/clockwise/core"
	"github.com/alecsavvy/clockwise/core/chain"
	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/ports/graph"
	"github.com/alecsavvy/clockwise/utils"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func run() error {
	// logger setup
	logger := utils.NewLogger(nil)
	logger.Info("good morning")

	// config setup
	homeDir := "./cmt-home"

	// db setup
	ctx := context.Background()
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

	db := db.New(pool)

	// chain setup
	node, err := chain.New(logger, homeDir, pool)
	if err != nil {
		return utils.AppError("failure to init chain", err)
	}

	clockwiseCore := core.NewCore(logger, node.Node(), db)

	// graphql setup
	gqlResolver := graph.NewResolver(logger, clockwiseCore)
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

	// web server setup
	e := echo.New()

	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/graphiql", graphiqlHandler)
	e.Any("/query", queryHandler)

	// run all the processes
	var wg sync.WaitGroup

	wg.Add(2)

	// run chain
	go func() {
		defer wg.Done()
		node.Run()
	}()

	// run web server
	go func() {
		defer wg.Done()
		err = e.Start(":26659")
		if err != nil {
			logger.Error("web server crashed", err)
		}
	}()

	// // run pubsub listener
	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	defer wg.Done()
	// 	if err := pubsub.Run(); err != nil {
	// 		logger.Error("pubsub crashed", err)
	// 	}
	// }()

	wg.Wait()
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println("fatal error: ", err)
	}

}
