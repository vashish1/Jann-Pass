package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/vashish1/Jann-Pass/utilities"
)

//TODO- same qr request -No increase in counter for area

type QRStr struct {
	Base64String string `json:"base_64_string,omitempty"`
}

func checkepass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	//Authenticating police code
	ok := utilities.AuthPoliceVerification(tokenString)
	if ok {
		var str QRStr
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &str)
		if err != nil {
			res := Response{
				Error: err.Error(),
			}
			b, _ := json.Marshal(res)
			w.Write(b)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ok := utilities.ValidateQR(str.Base64String)
		if ok {
			res := Response{
				Success: true,
				Error:   "nil",
			}
			b, _ := json.Marshal(res)
			w.Write(b)
			w.WriteHeader(http.StatusOK)
			return
		}

		res := Response{
			Success: false,
			Error:   "Invalid QR",
		}
		fmt.Println(res)
		b, _ := json.Marshal(res)
		w.Write(b)
		// json.NewEncoder(w).Encode(res)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//If user is unauthorized
	res := Response{
		Error: "User Authorization failed",
	}
	b, _ := json.Marshal(res)
	w.Write(b)
	w.WriteHeader(http.StatusBadRequest)
	return

}
