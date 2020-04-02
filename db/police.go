package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ID struct {
	ID int
	Token string
}

func InsertID(c *mongo.Collection, id int) bool {
	var x ID
	x.ID = id
	insertResult, err := c.InsertOne(context.TODO(), x)
	if err != nil {
		log.Print(err)
		return false
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return true
}

func ValidID(c *mongo.Collection, id int) bool {
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	var result ID

	err := c.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true

}

func UpdatePoliceCreds(c *mongo.Collection,id int,token string)bool{
	filter := bson.D{
		{"id", id},
	}
	update := bson.D{
		{
			"$set", bson.D{{"token", token}},
		},
	}
	updateResult, err := c.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return true
}