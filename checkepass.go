package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/vashish1/Jann-Pass/utilities"
)

//fix updateUSer
type QRStr struct {
	base64String string
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
				Error: err,
			}
			b, _ := json.Marshal(res)
			w.Write(b)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ok := utilities.ValidateQR(str.base64String)
		if ok {
			res := Response{
				Success: true,
				Error:   nil,
			}
			b, _ := json.Marshal(res)
			w.Write(b)
			w.WriteHeader(http.StatusOK)
			return
		}

		res := Response{
			Error:   errors.New("Invalid QR"),
			Success: false,
		}
		b, _ := json.Marshal(res)
		w.Write(b)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//If user is unauthorized
	res := Response{
		Error: errors.New("User Authorization failed"),
	}
	b, _ := json.Marshal(res)
	w.Write(b)
	w.WriteHeader(http.StatusBadRequest)
	return

}
