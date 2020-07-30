package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vashish1/Jann-Pass/db"
)

type mocksignup struct {
	Name     string
	Aadhar   string
	Email    string
	Password string
}

func signup(w http.ResponseWriter, r *http.Request) {
	var regis mocksignup
	w.Header().Set("Content-Type", "application/json")
	// var user db.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &regis)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}
	u := db.Newuser(regis.Name, regis.Email, regis.Aadhar, regis.Password)
	fmt.Println("user here", u)
	ok, er := db.Insertintouserdb(cl1, u)
	if ok {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"successful": "Registered"}`))
	} else {
		fmt.Println(er)
		json.NewEncoder(w).Encode(er)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "user not created"}`))
	}
}
