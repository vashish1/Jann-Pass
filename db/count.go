package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Epass stores the data of epass
type Area struct {
	Name  string `json:"name,omitempty"`
	City  string `json:"city,omitempty"`
	Code  int    `json:"code,omitempty"`
	count int    `json:"count,omitempty"`
}

//InsertEpass inserts the data into the database
func addArea(area Area) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	insertResult, err := countCl.InsertOne(ctx, area)
	if err != nil {
		log.Print(err)
		return false
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return true
}

//GetAreaList return the list of all the area present in databse.
func GetAreaList() ([]Area, error) {

	findOptions := options.Find()
	var result []Area

	cur, err := countCl.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {

	}
	for cur.Next(context.TODO()) {
		var elem Area
		err := cur.Decode(&elem)
		if err != nil {
		}
		result = append(result, elem)
	}
	if err := cur.Err(); err != nil {
		return []Area{}, err
	}

	cur.Close(context.TODO())
	return result, nil
}

//To get the count of Epass issued so far
func getCountforArea(code int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{primitive.E{Key: "code", Value: code}}
	var result Area

	err := countCl.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return 0, err
	}
	return result.count, nil
}

//To Reset all the counters after a week.
func ResetCounters() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter:=bson.M{}
    update:=bson.D{
		{"$set", bson.D{{"count", 0}}},
	}
    result, err := countCl.UpdateMany(
		ctx,
		filter,
		update,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
}

//To check if the count of epass issued is <50 (max count for epass in a week)
func CheckCounter(a Area) bool {

	// if the area already exists in database
	// then check the count, for epass eligiblity
	ok := areaExists(a.Code)
	if ok {
		count, err := getCountforArea(a.Code)
		if err != nil {
			return false
		}
		if count < 50 {
			return true
		}
		return false
	}

	// for new area, insert in db and return true
	//if not added return false
	if addArea(a) {
		return true
	}
	return false
}

//To check if the given area exists in database or not.
func areaExists(code int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{primitive.E{Key: "code", Value: code}}
	var result Area

	err := countCl.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return false
	}
	return true
}
