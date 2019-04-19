package database

import (
	"context"
	"fmt"
	"log"

	"../utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBInstance can be used in any service after connected
var DBInstance *mongo.Database

// Connect to the database
func Connect() *mongo.Database {
	// Check whether the .env variables exist
	dbName, err := utils.Env("DB_NAME")
	if err != nil {
		log.Fatal("DB_NAME not declared in .env file")
	}
	connectionString, err := utils.Env("DB_URL")
	if err != nil {
		log.Fatal("DB_URL not declared in .env file")
	}

	// Connect to the database
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	DBInstance = client.Database(dbName)
	return DBInstance
}
