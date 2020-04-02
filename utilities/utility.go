package utilities

import (
	// "fmt"
	"Jann-Pass/db"
	b64 "encoding/base64"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(name, email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":  name,
		"email": email,
	})

	tokenString, err := token.SignedString([]byte("idgafaboutthingsanymore"))
	if err == nil {
		return tokenString
	}
	return ""
}

func EncodeQrString(d db.Epass) string {

	data := d.Email + "," + d.Slot + "," + d.Date
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	return sEnc

}

func DecodeQrString(data string) string {

	sDec, _ := b64.StdEncoding.DecodeString(data)
	return string(sDec)
}

func StoreImage(str string) string {
	url := "https://api.qrserver.com/v1/create-qr-code/?data=[" + str + "]&size=[pixels]x[pixels]"
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("http.Get -> %v", err)
	}

	// We read all the bytes of the image
	// Types: data []byte
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("ioutil.ReadAll -> %v", err)
	}

	// You have to manually close the body, check docs
	// This is required if you want to use things like
	// Keep-Alive and other HTTP sorcery.
	res.Body.Close()
    path:="./Qr/"+str[0:5]+".png"
	// You can now save it to disk or whatever...
	ioutil.WriteFile(path, data, 0666)

	log.Println("I saved your image buddy!")
	return path
}
