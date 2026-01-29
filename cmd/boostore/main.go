package main

import (
	"context"
	"fmt"
	"log"

	"github.com/BlackHole55/software-store-final/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := config.ConnectDB()

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			fmt.Printf("Error closing DB: %v\n", err)
		}
	}()
}
