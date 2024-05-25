package main

import (
	"fmt"
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
	"github.com/alecsavvy/clockwise/graph"
	"github.com/alecsavvy/clockwise/utils"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func run() error {
	// logger setup
	logger := utils.NewLogger(nil)
	logger.Info("good morning")

	// config setup
	homeDir := "./cmt-home"

	coreApp := core.NewCore(logger)
	node, err := core.NewNode(logger, homeDir, coreApp)
	if err != nil {
		return utils.AppError("failure to init chain", err)
	}

	// graphql setup
	gqlResolver := graph.NewResolver(logger, coreApp)
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
		err = e.Start(":26659")
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
