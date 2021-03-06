package utilities

import (
	// "fmt"
	b64 "encoding/base64"
	"math/rand"
	"strings"
	"time"

	"github.com/vashish1/Jann-Pass/db"
)

//TODO: how to cross check if the user is using
// the pass in the same area as the one he issued it for.

//encoding data to string
func EncodeQrString(d db.Epass) string {

	var hash string
	var expire, start time.Time
	y, m, date := time.Now().Date()

	//Hash value according to the slot
	switch d.Slot {
	case "11:00-1:00":
		start = time.Date(y, m, date, 11, 0, 0, 0, time.Local)
		expire = time.Date(y, m, date, 13, 0, 0, 0, time.Local)
		hash = strings.Join([]string{d.Email, start.String(), expire.String()}, ",")
	case "2:00-4:00":
		start = time.Date(y, m, date, 13, 0, 0, 0, time.Local)
		expire = time.Date(y, m, date, 16, 0, 0, 0, time.Local)
		hash = strings.Join([]string{d.Email, start.String(), expire.String()}, ",")
	case "5:00-7:00":
		start = time.Date(y, m, date, 17, 0, 0, 0, time.Local)
		expire = time.Date(y, m, date, 19, 0, 0, 0, time.Local)
		hash = strings.Join([]string{d.Email, start.String(), expire.String()}, ",")
	}

	//Encoding the string to Base64 hash
	sEnc := b64.StdEncoding.EncodeToString([]byte(hash))
	return sEnc

}


//Decoding the QR encoded String
func DecodeQrString(data string) []string {
	sDec, _ := b64.StdEncoding.DecodeString(data)
	enc := strings.Split(string(sDec), ",")
	return enc
}

//Helper functio to generate ID's
func GenerateID() []int {
	var slice []int
	for i := 0; i < 20; i++ {
		rand.Seed(time.Now().UnixNano())
		min := 1000
		max := 3000
		slice = append(slice, (rand.Intn(max-min+1) + min))
	}
	return slice
}

//To Validate the QR encoded string 
func ValidateQR(encodedString string) bool {

	data := DecodeQrString(encodedString)
	db.ResetUser(data[0])
	layout := "2006-01-02 15:04:05 -0700 MST"
	start, _ := time.ParseInLocation(layout, data[1], time.Local)
	end, _ := time.ParseInLocation(layout, data[2], time.Local)
	now := time.Now()
	if start.Before(now) && end.After(now) {
		return true
	}
	return false
}

// <--------------------------------No Longer Required ----------------------------------->

// func StoreImage(str string) string {
// 	url := "https://api.qrserver.com/v1/create-qr-code/?data=[" + str + "]&size=[pixels]x[pixels]"
// 	res, err := http.Get(url)
// 	if err != nil {
// 		log.Fatalf("http.Get -> %v", err)
// 	}

// 	// We read all the bytes of the image
// 	// Types: data []byte
// 	data, err := ioutil.ReadAll(res.Body)

// 	if err != nil {
// 		log.Fatalf("ioutil.ReadAll -> %v", err)
// 	}

// 	// You have to manually close the body, check docs
// 	// This is required if you want to use things like
// 	// Keep-Alive and other HTTP sorcery.
// 	res.Body.Close()
// 	path := "./Qr/" + str[0:5] + ".png"
// 	// You can now save it to disk or whatever...
// 	ioutil.WriteFile(path, data, 0666)

// 	log.Println("I saved your image buddy!")
// 	return path
// }

//<-----------------------------No Longer Required ----------------------------------->
