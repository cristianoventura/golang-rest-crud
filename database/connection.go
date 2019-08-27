package database

import (
	"context"
	"fmt"
	"golang-rest-mongodb/structs"
	"golang-rest-mongodb/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	*mongo.Database
}

type ServiceInterface interface {
	AddPerson(person structs.Person) (interface{}, error)
	FindPerson(id string) (interface{}, error)
}

// Connect to the database
func Connect() (*DB, error) {
	// Check whether the .env variables exist
	dbName, err := utils.Env("DB_NAME")
	if err != nil {
		return nil, err
	}
	connectionString, err := utils.Env("DB_URL")
	if err != nil {
		return nil, err
	}

	// Connect to the database
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB")

	return &DB{
		client.Database(dbName),
	}, nil
}
