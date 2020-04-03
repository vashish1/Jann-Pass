package main

import (
	"Jann-Pass/db"
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	ok := db.Insertintouserdb(cl1, u)
	if ok {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"successful": "Registered"}`))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "user not created"}`))
	}
}