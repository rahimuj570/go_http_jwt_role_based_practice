package jwtconfig

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func PurseJWT(bearrer_token string, w http.ResponseWriter) (any, any) {
	token, err := jwt.Parse(bearrer_token, func(t *jwt.Token) (interface{}, error) {
		return Secret, nil
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err, err
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims["id"], claims["role"]
}
