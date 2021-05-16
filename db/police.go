package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ID struct {
	ID int
}

//TODO: Admin routes to add or delete police

//To add a new PoliceID in the database
func InsertID(id int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var x ID
	x.ID = id
	insertResult, err := policeCl.InsertOne(ctx, x)
	if err != nil {
		log.Print(err)
		return false
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return true
}

//To Validate the ID of the POLICE making a Request
func ValidID(id int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	var result ID

	err := policeCl.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true

}