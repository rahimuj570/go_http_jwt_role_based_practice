package main

import (
	"net/http"

	"github.com/rahimuj570/go_http_jwt_role_based_practice/controllers"
	"github.com/rahimuj570/go_http_jwt_role_based_practice/middlewares"
)

func main() {
	mux := http.NewServeMux()

	V_GetTodoForAdmin := controllers.GetTodoForAdmin

	mux.HandleFunc("GET /", controllers.Test)
	mux.HandleFunc("POST /", controllers.Login)
	mux.HandleFunc("GET /get", controllers.GetText)
	mux.Handle("GET /get_admin", middlewares.AdminMiddleware(http.HandlerFunc(V_GetTodoForAdmin)))
	mux.HandleFunc("GET /get_editor", controllers.GetTodoForEditor)

	http.ListenAndServe(":8080", mux)
}
