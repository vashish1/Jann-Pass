package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ID struct {
	ID    int
	Token string
}

func InsertID(c *mongo.Collection, id int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var x ID
	x.ID = id
	insertResult, err := c.InsertOne(ctx, x)
	if err != nil {
		log.Print(err)
		return false
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return true
}

func ValidID(c *mongo.Collection, id int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	var result ID

	err := c.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true

}

func UpdatePoliceCreds(c *mongo.Collection, id int, token string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{
		{"id", id},
	}
	update := bson.D{
		{
			"$set", bson.D{{"token", token}},
		},
	}
	updateResult, err := c.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return true
}
