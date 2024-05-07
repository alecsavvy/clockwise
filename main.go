package main

import (
	"log"
	"sync"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alecsavvy/clockwise/chain"
	"github.com/alecsavvy/clockwise/db"
	"github.com/alecsavvy/clockwise/graph"
	"github.com/alecsavvy/clockwise/utils"
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
	_, err := db.New()
	if err != nil {
		return utils.AppError("db initialization error", err)
	}

	// chain setup
	node, err := chain.New(homeDir)
	if err != nil {
		return utils.AppError("failure to init chain", err)
	}

	// graphql setup
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	queryHandler := func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	}

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
	e.POST("/query", queryHandler)

	// run all the processes
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		node.Run()
	}()

	go func() {
		defer wg.Done()
		err = e.Start(":26659")
		if err != nil {
			logger.Error("web server crashed", err)
		}
	}()

	wg.Wait()
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal("fatal error", err)
	}
}
