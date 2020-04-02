package main

import (
	"Jann-Pass/utilities"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type str struct {
	Qr string
}

func checkepass(w http.ResponseWriter, r *http.Request) {
	var test str
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &test)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}
	ok := utilities.IsQrValid(cl3, test.Qr)
	if ok {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Successfull": "Qr is valid"}`))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"error": "Qr is not valid"}`))
	}
}
