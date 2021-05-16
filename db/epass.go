package db

//Epass stores the data of epass
type Epass struct {
	Email    string `json:"email,omitempty"`
	Qr       string `json:"qr,omitempty"`
	Slot     string `json:"slot,omitempty"`
	Date     string `json:"date,omitempty"`
	Area     string `json:"area,omitempty"`
	AreaCode int    `json:"area_code,omitempty"`
}


// <------- To be implemented Later -------------->

// //InsertEpass inserts the data into the database
// func InsertEpass(u Epass) bool {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	insertResult, err := cl3.InsertOne(ctx, u)
// 	if err != nil {
// 		log.Print(err)
// 		return false
// 	}

// 	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
// 	return true
// }

// func EpassExists(c *mongo.Collection, email, enc string) bool {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	filter := bson.D{primitive.E{Key: "email", Value: email}, {Key: "qr", Value: enc}}
// 	var result Epass

// 	err := c.FindOne(ctx, filter).Decode(&result)
// 	if err != nil {
// 		fmt.Println(err)
// 		return false
// 	}
// 	return true
// }

// func DeleteEpass(c *mongo.Collection, email, enc string) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	deleteResult, err := c.DeleteOne(ctx, bson.D{primitive.E{Key: "email", Value: email}, {Key: "qr", Value: enc}})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
// }
