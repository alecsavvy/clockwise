package main

import (
	"log"

	"github.com/alecsavvy/clockwise/app"
	"github.com/alecsavvy/clockwise/common"
	"github.com/alecsavvy/clockwise/db"
	"github.com/alecsavvy/clockwise/discovery"
	"github.com/alecsavvy/clockwise/grpc"
)

func run() error {
	host := "localhost:6000"

	logger, err := common.NewLogger()
	if err != nil {
		return err
	}

	db, err := db.New()
	if err != nil {
		return err
	}

	discovery, err := discovery.New(host)
	if err != nil {
		return err
	}

	grpc, err := grpc.New(host)
	if err != nil {
		return err
	}

	app, err := app.New(logger, grpc, db, discovery)
	if err != nil {
		return err
	}

	return app.Run()
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("clockwise crashed: %s", err)
	}
}
