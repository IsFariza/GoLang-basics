package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var DB *mongo.Client

func ConnectDB() *mongo.Client {
	mongoURI := os.Getenv("MONGODB_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize MongoDB client: %v", err))
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(fmt.Sprintf("Failed to ping MongoDB: %v", err))
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	DB = client
	return client
}
