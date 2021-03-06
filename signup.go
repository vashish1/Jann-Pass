package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/vashish1/Jann-Pass/db"
)

type mocksignup struct {
	Name     string `json:"name,omitempty"`
	Aadhar   string `json:"aadhar,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type Response struct {
	Error   string `json:"error,omitempty"`
	Success bool   `json:"success,omitempty"`
}

func signup(w http.ResponseWriter, r *http.Request) {
	var regis mocksignup
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &regis)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}
	u := db.Newuser(regis.Name, regis.Email, regis.Aadhar, regis.Password)
	ok, err := db.Insertintouserdb(u)
	if ok {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": "true"}`))
	} else {
		res := Response{
			Error: err.Error(),
		}
		b, _ := json.Marshal(res)
		w.Write(b)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
