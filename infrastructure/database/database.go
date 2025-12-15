package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=123456 dbname=bookdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db
	fmt.Println("Database connected")
	return db
}
