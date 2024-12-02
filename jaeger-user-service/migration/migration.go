package main

import (
	"fmt"
	"jarger-user-service/helper"
	"jarger-user-service/models"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var PSQL_CONNECTION_USER string

func init() {
	var err error
	if err = godotenv.Load("../.env"); err != nil {
		panic(fmt.Sprintf("Error while loading .env: %s", err.Error()))
	}
	PSQL_CONNECTION_USER = helper.GetENV("PSQL_CONNECTION_USER", "")
}

func connectPSQL(conn string) *gorm.DB {
	var (
		connection *gorm.DB
		err        error
		gormLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		)
	)
	if connection, err = gorm.Open(postgres.Open(conn), &gorm.Config{
		Logger: gormLogger,
	}); err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %s", err.Error()))
	}
	return connection
}

func main() {
	db := connectPSQL(PSQL_CONNECTION_USER)
	db.Migrator().AutoMigrate(
		&models.User{},
	)
}
