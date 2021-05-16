package main

import (
	"encoding/json"
	"image/png"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/vashish1/Jann-Pass/db"
	"github.com/vashish1/Jann-Pass/utilities"
)

//TODO: Automatically modify the user data if he/she
//does not use the epass issued for that day, at the
//end of the day.

func epass(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "image/png")
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	//Authenticating user
	user := utilities.AuthVerification(tokenString)
	if user.Email != "" {
		//check if user has already initialised
		//an epass for that day.
		if !user.EpassIssued {
			db.UpdateUser(user.Email)
			var epass db.Epass
			body, _ := ioutil.ReadAll(r.Body)
			err := json.Unmarshal(body, &epass)
			if err != nil {
				res := Response{
					Error: err.Error(),
				}
				b, _ := json.Marshal(res)
				w.Write(b)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			epass.Email = user.Email
			encodedString := utilities.EncodeQrString(epass)

			//creating a QR code and then sending as response
			qrCode, _ := qr.Encode(encodedString, qr.L, qr.Auto)
			qrCode, _ = barcode.Scale(qrCode, 128, 128)
			png.Encode(w, qrCode)
			w.WriteHeader(http.StatusAccepted)
			return

		}

		//Return error if EPass already issued by the user
		//ONE-PASS-PER-USER-PER-DAY
		res := Response{
			Error: "Epass limit Exceeded",
		}
		b, _ := json.Marshal(res)
		w.Write(b)
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
