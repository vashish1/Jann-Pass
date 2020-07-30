package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Createdb creates a database
func Createdb() (*mongo.Collection, *mongo.Collection, *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := os.Getenv("DbURL")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		db,
	))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	user := client.Database("Jann-Pass").Collection("User")
	police := client.Database("Jann-Pass").Collection("Police")
	Qr := client.Database("Jann-Pass").Collection("Epass")

	return user, police, Qr
}
