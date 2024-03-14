package main

import (
	"context"
	"fmt"
	"github/andrewmkano/gorm-batch-insert/internal"
	"github/andrewmkano/gorm-batch-insert/internal/database"
	"log"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Benchmark - BatchSize = 200
// Batch Inserting with GORM 10,000 took 83.410708ms
// More to come...
//
// Benchmark - BatchSize = 500
// Batch Inserting with GORM 10,000 took 67.422667ms
// Batch Inserting with GORM 20,000 took 110.224666ms
// Batch Inserting with GORM 100,000 took 471.939583ms
// Batch Inserting with GORM 1,000,000 took 2.769831667s

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load env file: ", err.Error())
	}

	err = database.Setup()
	if err != nil {
		log.Fatalf("failed while setting up the database connection: %s", err.Error())
	}

	fmt.Println("Database started ✅")

	var in string
	fmt.Println("How many records would you like to insert? ")
	fmt.Scanln(&in)

	recordsToInsert, err := strconv.Atoi(in)
	if err != nil {
		log.Fatalf("invalid number of records: %s", err.Error())
	}

	ctx := context.Background()
	contacts := internal.GenerateDummyContacts(recordsToInsert)

	db := database.Conn.WithContext(ctx)

	// Just to get an idea of how long its taking us to insert that many records (roughly speaking)
	n := time.Now()
	db = db.Create(&contacts)
	if db.Error != nil {
		log.Fatalf("failed to insert contacts in batch: %s", err.Error())
	}

	m := message.NewPrinter(language.English)
	m.Printf("Batch Inserting with GORM %d took %s\n", recordsToInsert, time.Since(n))
}
