package main

import (
	"Jann-Pass/db"
	"Jann-Pass/utilities"
	"encoding/json"
	"fmt"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type pass struct {
	Aadhar string
	Slot   string
	Date   string
	Time   string
	Area string
}

func epass(w http.ResponseWriter, r *http.Request) {
	now:=time.Now()
	if count ==0{
		start=time.Now()
		finish=start.AddDate(0,0,7)
	}else if count==50&&now.Before(finish){
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "cannot issue pass: Limit exceeded for a week"}`))
		return
	}else if count==50&&now.After(finish){
		count=1
		start=now
		finish=start.AddDate(0,0,7)
	}else{
		count++
	}
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	fmt.Println("token", tokenString)

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("idgafaboutthingsanymore"), nil
	})
	var _, email string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		_ = claims["name"].(string)
		email = claims["email"].(string)
	}
	fmt.Print("email",email)
	is:=db.Findfromuserdb(cl1,email)
	if is{
		var regis pass
		w.Header().Set("Content-Type", "application/json")
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &regis)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "body not parsed"}`))
			return
		}
		var PASS db.Epass
		PASS.Email = email
		PASS.Aadhar = regis.Aadhar
		PASS.Slot = regis.Slot
		PASS.Date = regis.Date
		PASS.Area=regis.Area
		PASS.Qr = utilities.EncodeQrString(PASS)
		PASS.QrAddress = utilities.StoreImage(PASS.Qr)
		ok := db.InsertEpass(cl3, PASS)
		if ok {
			 var enc png.Encoder
			file, err := os.Open(PASS.QrAddress)
			if err != nil {
				fmt.Println(err)
			}
			 
			img, err := png.Decode(file)
			if err != nil {
				fmt.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"error": "Image not created"}`))
			}
	
			defer file.Close()
			w.Header().Set("Content-Type", "image/png")
			er:=enc.Encode(w, img)
			fmt.Print(er)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Pass not inserted"}`))
		}
	}else{
		w.WriteHeader(http.StatusOK)
	    w.Write([]byte(`{"error": "Authentication unsuccessful"}`))
	}
}
