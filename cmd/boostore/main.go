package main

import (
	"log"

	"github.com/BlackHole55/software-store-final/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()
}
