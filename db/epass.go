package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Epass stores the data of epass
type Epass struct {
	Email     string
	Qr        string
	QrAddress string
	Aadhar    string
	Slot      string
	Date      string
	Area      string
}

//InsertEpass inserts the data into the database
func InsertEpass(c *mongo.Collection, u Epass) bool {

	insertResult, err := c.InsertOne(context.TODO(), u)
	if err != nil {
		log.Print(err)
		return false
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return true
}

func EpassExists(c *mongo.Collection, email, enc string) bool {
	filter := bson.D{primitive.E{Key: "email", Value:email}, {Key:"qr",Value:  enc}}
	 var result Epass

	 err := c.FindOne(context.TODO(), filter).Decode(&result)
         if err != nil {
		  fmt.Println(err)
	       return false
			 }
	 return true
}

func DeleteEpass(c *mongo.Collection, email, enc string){

deleteResult, err := c.DeleteOne(context.TODO(),bson.D{primitive.E{Key: "email", Value:email}, {Key:"qr",Value:  enc}} )
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}