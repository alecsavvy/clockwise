package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func New() (*DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&User{}, &Track{})

	return &DB{
		db: db,
	}, nil
}
