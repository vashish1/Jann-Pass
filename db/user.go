package db

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User ......
type User struct {
	Name         string `json:"name,omitempty"`
	Email        string `json:"email,omitempty"`
	Aadhar       string `json:"aadhar,omitempty"`
	PasswordHash string `json:"password_hash,omitempty"`
	EpassIssued  bool   `json:"epass_issued,omitempty"`
}

//Newuser .....
func Newuser(name, email, aadhar, password string) User {

	Password := SHA256ofstring(password)
	U := User{Name: name, Email: email, PasswordHash: Password, Aadhar: aadhar,EpassIssued: false}
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
func Insertintouserdb(u User) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	insertResult, err := userCl.InsertOne(ctx, u)
	if err != nil {
		log.Print("error in inserting user", err)
		return false, err
	}

	fmt.Println("Inserted a user document: ", insertResult.InsertedID)
	return true, nil
}

//Findfromuserdb finds the required data
func UserExists(st string) bool {
	filter := bson.D{primitive.E{Key: "email", Value: st}}
	var result User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := userCl.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//FindUser finds if the user exists but with respect to the username.
func FindUser(st string, p string) (bool,User) {
	filter := bson.D{primitive.E{Key: "email", Value: st}}
	var result User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := userCl.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return false,User{}
	}
	if result.PasswordHash != SHA256ofstring(p) {
		return false,User{}
	}
	return true,result
}

//Finddb returs the user detail with respect to email string
func Finddb(st string) (User,error) {
	filter := bson.D{primitive.E{Key: "email", Value: st}}
	var result User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := userCl.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return User{},err
	}

	return result,nil
}


// UpdateUser updates the user info accordingly
func UpdateUser(email string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{
		{"email", email},
	}
	update := bson.D{
		{
			"$set", bson.D{{"epass_issued", true}},
		},
	}
	updateResult, err := userCl.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return true
}

//ResetUser unset the pass-issued value
func ResetUser(email string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{
		{"email", email},
	}
	update := bson.D{
		{
			"$set", bson.D{{"epass_issued", false}},
		},
	}
	updateResult, err := userCl.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return true
}
