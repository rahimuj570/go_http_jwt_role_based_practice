package jwtconfig

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func PurseJWT(bearrer_token string, w http.ResponseWriter) any {
	token, err := jwt.Parse(bearrer_token, func(t *jwt.Token) (interface{}, error) {
		return Secret, nil
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	claims := token.Claims
	return claims
}
