package internal

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var BatchSize int

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Failed to load env file: ", err.Error())
		os.Exit(1)
	}

	bsize := os.Getenv("BATCH_SIZE")
	BatchSize, err = strconv.Atoi(bsize)
	if err != nil {
		fmt.Println("failed to convert batch size env var into an integer: ", err.Error())
		os.Exit(1)
	}
}
