package main

import (
	"Jann-Pass/db"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type police struct {
	ID int
}

func policeLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// fmt.Fprintf(w, "Hello World! %s", time.Now())
	// var result database.User
	var user police
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}
	ok := db.ValidID(cl1, user.ID)
	if ok {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id": user.ID,
		})

		tokenString, err := token.SignedString([]byte("idgafaboutthingsanymore"))

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "error in token string"}`))
			return
		}
		var try t
		try.Success="Logged in successfully"
		try.Token = tokenString
		tkn := db.UpdatePoliceCreds(cl2, user.ID, tokenString)
		if tkn {
			json.NewEncoder(w).Encode(try)
			w.WriteHeader(http.StatusCreated)
			// w.Write([]byte(`{"success": "created token successfully"}`))
		} else {
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"error": "token not created"}`))
		}
	}else{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid Id"}`))
	}
}
