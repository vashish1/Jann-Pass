package main

import (
	"Jann-Pass/db"
	"time"
	// "Jann-Pass/utilities"

	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var cl1, cl2,cl3  *mongo.Collection
var count int
var start time.Time
var finish time.Time

func init(){
 cl1,cl2,cl3=db.Createdb()
 }

// 
func main() {
	r := mux.NewRouter()
	headers:=handlers.AllowedHeaders([]string{"Accept", "Content-Type","Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	methods:=handlers.AllowedMethods([]string{"POST","GET","PUT","DELETE"})
	origins:=handlers.AllowedOrigins([]string{"*"})
	r.HandleFunc("/signup",signup).Methods("POST")
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/epass",epass).Methods("GET","POST")
	r.HandleFunc("/checkepass",checkepass).Methods("GET","POST")
	r.HandleFunc("/login/police",policeLogin).Methods("POST")
	http.Handle("/", handlers.CORS(headers,methods,origins)(r))
	http.ListenAndServe(":3000", nil)
}