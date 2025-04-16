package main

import (
	"net/http"

	"github.com/rahimuj570/go_http_jwt_role_based_practice/controllers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", controllers.Test)
	mux.HandleFunc("GET /get", controllers.GetText)

	http.ListenAndServe(":8080", mux)
}
