package db

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//User ......
type User struct {
	Name         string
	Email        string
	Aadhar       string
	PasswordHash string
	Token        string
}

//Newuser .....
func Newuser(name, email,aadhar, password string) User {

	Password := SHA256ofstring(password)
	U := User{Name: name, Email: email, PasswordHash: Password,Aadhar: aadhar}
	return U
}

//SHA256ofstring is a function which takes a string a returns its sha256 hashed form
func SHA256ofstring(p string) string {
	h := sha1.New()
	h.Write([]byte(p))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}

//Insertintouserdb inserts the data into the database
func Insertintouserdb(usercollection *mongo.Collection, u User) bool {

	fmt.Println(u.Name)
	insertResult, err := usercollection.InsertOne(context.TODO(), u)
	if err != nil {
		log.Print(err)
		return false
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return true
}

//Findfromuserdb finds the required data
func Findfromuserdb(usercollection *mongo.Collection, st string) bool {
	filter := bson.D{primitive.E{Key: "email", Value: st}}
	var result User

	err := usercollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//FindUser finds if the user exists but with respect to the username.
func FindUser(usercollection *mongo.Collection, st string, p string) bool {
	filter := bson.D{primitive.E{Key: "email", Value: st}}
	var result User

	err := usercollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if result.PasswordHash != SHA256ofstring(p) {
		return false
	}
	return true
}

func Finddb(usercollection *mongo.Collection, st string) User {
	filter := bson.D{primitive.E{Key: "email", Value: st}}
	var result User

	err := usercollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return result
	}
	
	return result
}


//UpdateToken updates the user info
func UpdateToken(c *mongo.Collection, o string, t string) bool {
	filter := bson.D{
		{"email", o},
	}
	update := bson.D{
		{
			"$set", bson.D{{"token", t}},
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

