package database

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Setup() error {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	bsize := os.Getenv("BATCH_SIZE")
	batchSize, err := strconv.Atoi(bsize)
	if err != nil {
		return fmt.Errorf("failed to convert batch size env var into an integer: %w", err)
	}

	Conn, err = gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort),
	}), &gorm.Config{
		CreateBatchSize: batchSize,
	})

	if err != nil {
		return fmt.Errorf("Could not open a connection to the database: %s", err.Error())
	}

	return nil
}
