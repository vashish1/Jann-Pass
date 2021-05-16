package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/vashish1/Jann-Pass/db"
)

type police struct {
	ID int
}

func policeLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ln police
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &ln)
	if err != nil {
		w.Write([]byte(`{"error": "body not parsed"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Check if the user exissts with such credentials
	ok := db.ValidID(ln.ID)
	if ok {

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id": ln.ID,
		})

		//generate Token-JWT
		tokenString, err := token.SignedString([]byte(os.Getenv("secret key")))
		if err != nil {
			res := Response{
				Error: err.Error(),
			}
			b, _ := json.Marshal(res)
			w.Write(b)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//Send a Successfull Response
		try := loginResponse{
			Success: "Log In successful",
			Token:   tokenString,
		}
		json.NewEncoder(w).Encode(try)
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{"error": "no such user exist"}`))
	return
}
