package utilities

import (
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(name,email string)string{
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":  name,
		"email": email,
	})

	tokenString, err := token.SignedString([]byte("idgafaboutthingsanymore"))
	if err==nil {
		return tokenString
	}
	return ""
}