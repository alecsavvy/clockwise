package main

import (
	"context"
	"log"

	"github.com/alecsavvy/clockwise/db"
)

func run() error {
	db, err := db.New()
	if err != nil {
		return err
	}

	_, err = db.Queries.ListAuthors(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("clockwise crashed %s", err)
	}
}
