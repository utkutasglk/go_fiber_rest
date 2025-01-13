package database

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){


	db, err := gorm.Open(sqlite.Open("myDatabase.db"))

	if err != nil {
		fmt.Println("Database connection error")
	}

	DB = db

	fmt.Println("Connected to database")
}