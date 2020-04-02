package main

import (
	"Jann-Pass/db"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"go.mongodb.org/mongo-driver/mongo"
)

var cl1, cl2  *mongo.Collection

func init(){
 cl1,cl2=db.Createdb()
}

// 
func main() {
	r := mux.NewRouter()
	headers:=handlers.AllowedHeaders([]string{"Accept", "Content-Type","Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	methods:=handlers.AllowedMethods([]string{"POST","GET","PUT","DELETE"})
	origins:=handlers.AllowedOrigins([]string{"*"})
	r.HandleFunc("/signup",signup).Methods("POST","GET")
	r.HandleFunc("/login", login).Methods("POST")
	http.Handle("/", handlers.CORS(headers,methods,origins)(r))
	http.ListenAndServe(":3000", nil)
}