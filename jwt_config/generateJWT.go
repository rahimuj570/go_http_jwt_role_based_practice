package jwtconfig

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rahimuj570/go_http_jwt_role_based_practice/enums"
)

var Secret = []byte("this is dummy secret")

func GenerateJWT(payload any, role enums.Role, w http.ResponseWriter) (string, error) {
	claims := jwt.MapClaims{
		"id":   payload,
		"role": role,
		"exp":  time.Now().Add(time.Minute * 3).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(Secret)
}
