package db

import (
	"context"
	"fmt"
	"log"
	"time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Createdb creates a database
func Createdb() (*mongo.Collection, *mongo.Collection, *mongo.Collection, context.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client, err := mongo.Connect(ctx, options.Client().ApplyURI(
  "mongodb+srv://yashi:KNT8CQsC7120GowL@cluster0-2pscc.mongodb.net/test?retryWrites=true&w=majority",
))
if err != nil { log.Fatal(err) }
	fmt.Println("Connected to MongoDB!")
	user := client.Database("Jann-Pass").Collection("User")
	police := client.Database("Jann-Pass").Collection("Police")
	Qr := client.Database("Jann-Pass").Collection("Epass")
	episodeResult, err := user.InsertMany(ctx, []interface{}{
		bson.D{
		
			{"title", "GraphQL for API Development"},
			{"description", "Learn about GraphQL from the co-creator of GraphQL, Lee Byron."},
			{"duration", 25},
		},
		bson.D{
			
			{"title", "Progressive Web Application Development"},
			{"description", "Learn about PWA development with Tara Manicsic."},
			{"duration", 32},
		},
	})
	if err != nil {
		log.Fatal("here",err)
	}
	fmt.Printf("Inserted %v documents into episode collection!\n", len(episodeResult.InsertedIDs))
	return user, police, Qr, ctx
}

// func SetContext() (context.Context,context.CancelFunc){
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	return ctx,cancel
// }