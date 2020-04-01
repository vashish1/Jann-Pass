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

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/signup",signup).Methods("POST","GET")
	r.HandleFunc("/login", login).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}