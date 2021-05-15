package utilities

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/vashish1/Jann-Pass/db"
)

var Secret = []byte(os.Getenv("secret key"))

func GenerateAuthToken(name, email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":  name,
		"email": email,
	})

	tokenString, err := token.SignedString([]byte("idgafaboutthingsanymore"))
	if err == nil {
		return tokenString
	}
	return ""
}

func AuthVerification(tokenString string) db.User {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return Secret, nil
	})

	var _, email string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		_ = claims["name"].(string)
		email = claims["email"].(string)
	}

	user, err := db.Finddb(email)
	if err != nil {
		return db.User{}
	}
	return user
}

func AuthPoliceVerification(tokenString string) bool {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return Secret, nil
	})

	var id int
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id = claims["id"].(int)
	}

	ok := db.ValidID(id)
	return ok
}
