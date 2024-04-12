package gorman

import (
	"fmt"
	"github/andrewmkano/gorm-batch-insert/internal"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type gormBatchManager struct {
	Conn *gorm.DB
}

func NewGormBatchManager() (gormBatchManager, error) {
	conn, err := setup()
	if err != nil {
		return gormBatchManager{}, fmt.Errorf("failed to start a gorm batch manager: %w", err)
	}

	return gormBatchManager{conn}, nil
}

func setup() (*gorm.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	Conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort),
	}), &gorm.Config{
		CreateBatchSize: internal.BatchSize,
	})

	if err != nil {
		return nil, fmt.Errorf("Could not open a connection to the database: %s", err.Error())
	}

	return Conn, nil
}

func (gbm gormBatchManager) SaveContactsInBatches(c any) error {
	if gbm.Conn.Create(c); gbm.Conn.Error != nil {
		return fmt.Errorf("gorm: failed to insert records in batches: %w", gbm.Conn.Error)
	}

	return nil
}
