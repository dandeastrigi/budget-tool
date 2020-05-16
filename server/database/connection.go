package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TestDatabaseService : test database connection
func TestDatabaseService() {
	uri := os.Getenv("DATABASE_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Service - OK")
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}

// Transaction
func Transaction() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:123@localhost:27017"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
