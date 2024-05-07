package main

import (
	"log"
	"sync"

	"github.com/alecsavvy/clockwise/chain"
	"github.com/alecsavvy/clockwise/db"
	"github.com/alecsavvy/clockwise/utils"
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

	// run all the processes
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		node.Run()
	}()

	go func() {

	}()

	wg.Wait()
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal("fatal error", err)
	}
}
