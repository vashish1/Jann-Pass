package main

import (
	"Jann-Pass/db"

	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var cl1, cl2  *mongo.Collection

func init(){
 cl1,cl2=db.Createdb()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/signup",signup).Methods("POST","GET")
	r.HandleFunc("/login", login).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":80", nil)
}