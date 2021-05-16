package main

import (
	"fmt"
	"os"
	"time"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/vashish1/Jann-Pass/db"
)

var port = os.Getenv("PORT")
var ticker = time.NewTicker(time.Now().AddDate(0, 0, 7).Sub(time.Now()))

//To Reset the counter after every 7 days,
//this function will wait for a ticker, which 
//will then make a call to Reset function.
func set() {
	for {
		select {
		case t := <-ticker.C:
			db.ResetCounters()
			fmt.Println("Counter Reset Time: ", t)
		}

	}

}

func main() {
    
	//Goroutine to Reset counter
	go set()

	r := mux.NewRouter()
	//handling CORS externally
	headers := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	//Routes
	r.HandleFunc("/signup", signup).Methods("POST")
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/epass", epass).Methods("GET", "POST")
	r.HandleFunc("/checkepass", checkepass).Methods("GET", "POST")
	r.HandleFunc("/login/police", policeLogin).Methods("POST")

	//making routes.
	http.Handle("/", handlers.CORS(headers, methods, origins)(r))
	http.ListenAndServe(":"+port, nil)
}
