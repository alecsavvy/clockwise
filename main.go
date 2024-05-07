package main

import (
	"log"

	"github.com/alecsavvy/clockwise/chain"
	"github.com/alecsavvy/clockwise/db"
	"github.com/alecsavvy/clockwise/utils"
	"github.com/labstack/echo/v4"
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

	// web server setup
	e := echo.New()
	e.HideBanner = true

	// chain setup
	node, err := chain.New(homeDir)
	if err != nil {
		return utils.AppError("failure to init chain", err)
	}
	node.Run()
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal("fatal error", err)
	}
}
