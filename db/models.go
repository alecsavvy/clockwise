package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Handle string `gorm:"unique"`
	Bio    string
}

type Track struct {
	gorm.Model
	Title       string
	Description string
	UserID      uint
}
