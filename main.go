package main

import (
	"fmt"
	"os"
	"time"

	// "github.com/vashish1/Jann-Pass/utilities"

	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/vashish1/Jann-Pass/db"
	// "gopkg.in/mgo.v2/bson"
)

var port = os.Getenv("PORT")
var starttime=time.Now()
var RestartAt=starttime.AddDate(0,0,7)

func RunTimer(){
for{
 fmt.Println("Run Timer working")
 current:=time.Now()
	if current==RestartAt{
		current=RestartAt
		RestartAt.AddDate(0,0,7)
		db.ResetCounters()
	}
 }
}

func main() {

	// go RunTimer()

	r := mux.NewRouter()

	headers := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	r.HandleFunc("/signup", signup).Methods("POST")
	r.HandleFunc("/login", login).Methods("POST")
	// r.HandleFunc("/logout",logout).Methods("GET,POST")
	r.HandleFunc("/epass", epass).Methods("GET", "POST")
	r.HandleFunc("/checkepass", checkepass).Methods("GET", "POST")
	r.HandleFunc("/login/police", policeLogin).Methods("POST")
	http.Handle("/", handlers.CORS(headers, methods, origins)(r))
	http.ListenAndServe(":"+port, nil)
}

