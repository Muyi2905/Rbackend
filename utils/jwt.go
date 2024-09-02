package utils

import (
	"go/token"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(UserId string) (string, error) {
	claims:= jwt.Claims{
		"userId" : UserId,
		"exp" :    time.Now().Add(time.Hour*24).Unix()
	}

	token:= jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}