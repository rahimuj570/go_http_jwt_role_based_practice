package controllers

import (
	"encoding/json"
	"net/http"

	model "github.com/rahimuj570/go_http_jwt_role_based_practice/Model"
)

func GetText(w http.ResponseWriter, r *http.Request) {
	t := model.Todo{}
	json.NewEncoder(w).Encode(t.GetTodo())
}

func GetTodoForAdmin(w http.ResponseWriter, r *http.Request) {
	t := model.Todo{}
	json.NewEncoder(w).Encode(t.GetTodoForAdmin())
}

func GetTodoForEditor(w http.ResponseWriter, r *http.Request) {
	t := model.Todo{}
	json.NewEncoder(w).Encode(t.GetTodoForEditor())
}
