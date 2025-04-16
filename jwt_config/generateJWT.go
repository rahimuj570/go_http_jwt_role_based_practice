package jwtconfig

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var Secret = []byte("this is dummy secret")

func GenerateJWT(payload string, w http.ResponseWriter) (string, error) {
	claims := jwt.MapClaims{
		"data": payload,
		"exp":  time.Now().Add(time.Minute * 3).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(Secret)
}
