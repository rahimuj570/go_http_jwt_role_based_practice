package main

import (
	"net/http"

	"github.com/rahimuj570/go_http_jwt_role_based_practice/controllers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", controllers.Test)
	mux.HandleFunc("POST /", controllers.Login)
	mux.HandleFunc("GET /get", controllers.GetText)
	mux.HandleFunc("GET /get_admin", controllers.GetTodoForAdmin)
	mux.HandleFunc("GET /get_editor", controllers.GetTodoForEditor)

	http.ListenAndServe(":8080", mux)
}
