package main

import (
    "Jann-Pass/db"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type logn struct {
	Email string
	Password string
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var result database.User
	var user logn
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}
	ok := db.FindUser(cl1, user.Email, user.Password)
	if ok {
		u := db.Finddb(cl1, user.Email)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": u.Email,
	         "name": u.Name,
		})

		tokenString, err := token.SignedString([]byte("idgafaboutthingsanymore"))

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "error in token string"}`))
			return
		}
		type t struct{
			Token string
		}
		 var try t
		 try.Token=tokenString
		tkn := db.UpdateToken(cl1, u.Email, tokenString)
		if tkn {
			json.NewEncoder(w).Encode(try)
			w.WriteHeader(http.StatusCreated)
			// w.Write([]byte(`{"success": "created token successfully"}`))
		} else {
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"error": "token not created"}`))
		}
	}
}
