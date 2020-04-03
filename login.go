package main

import (
	"Jann-Pass/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type logn struct {
	Email    string
	Password string
}

func login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// fmt.Fprintf(w, "Hello World! %s", time.Now())
	// var result database.User
	var user logn
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}
	fmt.Println(user)
	ok := db.FindUser(cl1, user.Email, user.Password)
	if ok {
		fmt.Println("1")
		u := db.Finddb(cl1, user.Email)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": u.Email,
			"name":  u.Name,
		})

		tokenString, err := token.SignedString([]byte("idgafaboutthingsanymore"))
		fmt.Println("2")
		if err != nil {

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "error in token string"}`))
			return
		}
		type t struct {
			Token string
		}
		fmt.Println("3")
		var try t
		try.Token = tokenString
		tkn := db.UpdateToken(cl1, u.Email, tokenString)
		if tkn {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(try)
			// w.Write([]byte(`{"successful": "Registered"}`))
			// fmt.Print("try",try)
			// w.Write([]byte(`
			// {"success": "created token successfully"}
			// {"success":}`))
			// json.NewEncoder(w).Encode(try)
			// w.WriteHeader(http.StatusCreated)
			
		} else {
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"error": "token not created"}`))
		}
	} else {
		fmt.Print("4")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"error": "no such user exist"}`))
	}
}
