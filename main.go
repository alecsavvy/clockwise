package main

import (
	"flag"
	"log"

	"github.com/alecsavvy/clockwise/common"
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

	app, err := NewApp(logger, config)
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
