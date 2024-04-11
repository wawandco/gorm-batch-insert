package gsqlman

import (
	"database/sql"
	"fmt"
	"github/andrewmkano/gorm-batch-insert/internal"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

type goSqlBatchManager struct {
	Conn *sql.DB
}

func NewGoSQLBatchManager() (goSqlBatchManager, error) {
	conn, err := setup()
	if err != nil {
		return goSqlBatchManager{}, fmt.Errorf("failed to start a gosql batch manager: %w", err)
	}

	return goSqlBatchManager{conn}, nil
}

func setup() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName)

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to start database with sql package: %w", err)
	}

	return conn, nil
}

func (gsqbm goSqlBatchManager) SaveContactsInBatches(contacts internal.Contacts) error {
	contactGroups := contacts.SplitInGroups(internal.BatchSize)
	for _, contacsG := range contactGroups {
		var (
			placeholders []string
			vals         []interface{}
		)

		for index, contact := range contacsG {
			placeholders = append(placeholders, fmt.Sprintf("($%d,$%d,$%d)",
				index*3+1,
				index*3+2,
				index*3+3,
			))

			vals = append(vals, contact.FirstName, contact.LastName, contact.Email)
		}

		txn, err := gsqbm.Conn.Begin()
		if err != nil {
			return fmt.Errorf("could not start a new transaction: %w", err)
		}

		insertStatement := fmt.Sprintf("INSERT INTO contacts(first_name,last_name,email) VALUES %s", strings.Join(placeholders, ","))
		_, err = txn.Exec(insertStatement, vals...)
		if err != nil {
			txn.Rollback()
			return fmt.Errorf("failed to insert multiple records at once: %w", err)
		}

		if err := txn.Commit(); err != nil {
			return fmt.Errorf("failed to commit transaction: %w", err)
		}
	}

	return nil
}
