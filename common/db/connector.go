package db

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"ssff.com.br/openway/common/entity"
)

const dsn = "host=localhost port=5432 user=postgres password=postgres dbname=openway sslmode=disable"

var db *gorm.DB

func GetConnection() *gorm.DB {
	if db != nil {
		return db
	}

	logrus.Info("Connecting to database...")

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal("failed to connect database")
	}

	return db
}

func InitDB() {
	logrus.Info("Initializing database...")

	db := GetConnection()
	db.AutoMigrate(
		&entity.Api{},
		&entity.Endpoint{},
	)
}
