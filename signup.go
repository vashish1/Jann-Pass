package main

import (
	"Jann-Pass/db"
	"Jann-Pass/utilities"
	"net/http"
)

type mocksignup struct{
	Name string
	Aadhar string
	Email string
	Password string
}

func signup(w http.ResponseWriter,r *http.Request){
	var regis mocksignup
	w.Header().Set("Content-Type", "application/json")
	var user database.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &regis)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}
	u:=db.Newuser(mocksignup.Name,mocksignup.Email,mocksignup.Aadhar,mocksignup.Password)
	// token:=utilities.GenerateToken(u.Name,u.Email)
	// u.Token=db.UpdateToken(cl1,u.Email,token)
	ok:= db.Insertintouserdb(cl1,u)
	if ok{
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"successful": "Registered"}`))
	}else{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "user not created"}`))
	}
}