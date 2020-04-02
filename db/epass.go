package db

import (
	"context"
	"fmt"
	"log"

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
