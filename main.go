package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alecsavvy/clockwise/chain"
	"github.com/alecsavvy/clockwise/db"
	"github.com/alecsavvy/clockwise/ports/graph"
	"github.com/alecsavvy/clockwise/utils"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
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
	_ = db.New(pool)

	// chain setup
	node, err := chain.New(logger, homeDir)
	if err != nil {
		return utils.AppError("failure to init chain", err)
	}

	// rpc client setup
	rpcUrl := node.RPC()
	client, err := rpchttp.New(rpcUrl, "/websocket")
	if err != nil {
		return utils.AppError("failure to init chain rpc", err)
	}
	client.SetLogger(logger)

	// graphql setup
	gqlResolver := graph.NewResolver(logger, client)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: gqlResolver}))
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
		fmt.Println("fatal error: ", err)
	}

}
