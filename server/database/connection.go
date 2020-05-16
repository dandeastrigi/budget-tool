package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"server/model"
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

// SaveBudget : budget sav
func SaveBudget(budget model.Budget) {
	clientOptions := options.Client().ApplyURI(os.Getenv("DATABASE_URI"))
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("budget").Collection("budget")

	insertResult, err := collection.InsertOne(context.TODO(), budget)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
