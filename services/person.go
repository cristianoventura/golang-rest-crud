package services

import (
	"context"

	"../structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// AddPerson - Creates a new document in people collection
func AddPerson(person structs.Person, db *mongo.Database) (interface{}, error) {
	collection := db.Collection("people")
	res, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// FindPerson - Returns the document by the specified id
func FindPerson(id string, db *mongo.Database) (interface{}, error) {
	collection := db.Collection("people")
	filter := bson.D{{"_id", id}}
	var res structs.Person
	err := collection.FindOne(context.Background(), filter).Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
