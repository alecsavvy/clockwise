package main

import (
	"log"

	"github.com/alecsavvy/clockwise/app"
	"github.com/alecsavvy/clockwise/common"
	"github.com/alecsavvy/clockwise/db"
	"github.com/alecsavvy/clockwise/discovery"
)

func run() error {
	logger, err := common.NewLogger()
	if err != nil {
		return err
	}

	db, err := db.New()
	if err != nil {
		return err
	}

	discovery, err := discovery.New("http://localhost:6000")
	if err != nil {
		return err
	}

	app, err := app.New(logger, db, discovery)
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
