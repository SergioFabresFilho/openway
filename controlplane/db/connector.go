package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn = "host=localhost port=5432 user=postgres password=postgres dbname=openway sslmode=disable"

var db *gorm.DB

func GetConnection() *gorm.DB {
	if db != nil {
		return db
	}

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	return db
}
