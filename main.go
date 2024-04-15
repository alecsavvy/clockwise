package main

import (
	"flag"
	"log"

	"github.com/alecsavvy/clockwise/app"
	"github.com/alecsavvy/clockwise/common"
	"github.com/alecsavvy/clockwise/peer"
	"github.com/alecsavvy/clockwise/server"
	"github.com/alecsavvy/clockwise/storage"
)

func run() error {
	logger, err := common.NewLogger()
	if err != nil {
		return err
	}

	configPath := flag.String("config", "./clockwise.toml", "path to configuration file")
	flag.Parse()

	config, err := common.ReadConfig(*configPath)
	if err != nil {
		return err
	}

	db, err := storage.New()
	if err != nil {
		return err
	}

	discovery, err := peer.New(config)
	if err != nil {
		return err
	}

	server, err := server.New(config)
	if err != nil {
		return err
	}

	app, err := app.New(logger, server, db, discovery)
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
