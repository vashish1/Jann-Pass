package main

import (
	"Jann-Pass/db"
	"Jann-Pass/utilities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type str struct {
	Qr string
}

func checkepass(w http.ResponseWriter, r *http.Request) {
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
	var id int
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id = claims["id"].(int)
	}
	is := db.ValidID(cl2, id)
	if is {
		var test str
		w.Header().Set("Content-Type", "application/json")
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &test)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "body not parsed"}`))
			return
		}
		ok := utilities.IsQrValid(cl3, test.Qr)
		if ok {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"Successfull": "Qr is valid"}`))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"error": "Qr is not valid"}`))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Authentication unsuccessful"}`))
	}
}
