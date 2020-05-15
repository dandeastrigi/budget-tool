package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Budget struct {
	Total       float64
	Installment uint
}

// SaveBudget : save new budget
func SaveBudget(budget Budget) {
	clientOptions := options.Client().ApplyURI("mongodb://root:123@localhost:27017")
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

// FindBudget : find a budget
func FindBudget() {
	clientOptions := options.Client().ApplyURI("mongodb://root:123@localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("budget").Collection("budget")
	result, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for result.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		err := result.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result...

		// To get the raw bson bytes use cursor.Current
		raw := result.Current
		// do something with raw...
		log.Println(raw)
	}
	if err := result.Err(); err != nil {
		log.Println(err)
	}
}

// TestConnection : A new mongo conn
func TestConnection() {
	fmt.Println("Checking database!")
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://root:123@localhost:27017")

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
	fmt.Println("Database is Up!")

	err = client.Disconnect(context.TODO())
}
