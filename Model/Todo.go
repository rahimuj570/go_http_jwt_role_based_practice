package model

import "time"

//ToDo Struct
type Todo struct {
	Todo string    `json:"todo"`
	Time time.Time `json:"time"`
}

//Get Dummy ToDo Method
func (t *Todo) GetTodo() Todo {
	t.Time = time.Now()
	t.Todo = "Dummy Todo"
	return *t
}

func (t *Todo) GetTodoForAdmin() Todo {
	t.Time = time.Now()
	t.Todo = "Todo For Admin"
	return *t
}

func (t *Todo) GetTodoForEditor() Todo {
	t.Time = time.Now()
	t.Todo = "Todo For Editor"
	return *t
}
