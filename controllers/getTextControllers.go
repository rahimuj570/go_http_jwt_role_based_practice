package controllers

import (
	"encoding/json"
	"net/http"

	model "github.com/rahimuj570/go_http_jwt_role_based_practice/Model"
	jwtconfig "github.com/rahimuj570/go_http_jwt_role_based_practice/jwt_config"
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

func Login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	//
	//suppose we check the credential in DB// then👇
	//
	str_token, err := jwtconfig.GenerateJWT(user.Id, user.Role, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(struct{ Token string }{Token: str_token})

}
