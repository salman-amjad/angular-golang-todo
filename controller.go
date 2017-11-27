package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/gorilla/mux"
)

// Return Return object for JSON requests
type Return struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Todos   interface{} `json:"todos"`
}

// Index Index func
func Index(c *routing.Context) {
	c.Write("Welcome!\n")
}

// TodoIndex ewewe
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	res := Response{w}
	res.SendJSON(todos)
}

// TodoShow Show a single item by ID
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res := Response{w}
	var todoID int
	var err error
	_return := &Return{}
	if todoID, err = strconv.Atoi(vars["todoId"]); err != nil {
		panic(err)
	}
	todo := ModelFindTodo(todoID)
	if todo.Id > 0 {
		res.SendJSON(todo)
		return
	}

	_return.Success = false
	_return.Message = "Not found"
	res.SendJSON(_return)
}

// TodoDelete Delete an item by ID
func TodoDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res := Response{w}
	_return := &Return{}
	var todoID int
	var err error
	if todoID, err = strconv.Atoi(vars["todoId"]); err != nil {
		panic(err)
	}
	todo := ModelDestroyTodo(todoID)
	if todo == nil {
		_return.Success = true
		_return.Message = "Item removed successfully"
		_return.Todos = todos
		res.SendJSON(_return)
		return
	}

	_return.Success = false
	_return.Message = "Item not found"
	res.SendJSON(_return)
}

// TodoCreate Create an item using POST req
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	res := Response{w}
	_return := &Return{}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		panic(err)
	}

	ModelCreateTodo(todo)

	_return.Success = true
	_return.Message = "Item added successfully"
	_return.Todos = todos
	res.SendJSON(_return)
}

// TodoUpdate Create an item using POST req
func TodoUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todo Todo
	res := Response{w}
	_return := &Return{}
	var todoID int
	var err error
	if todoID, err = strconv.Atoi(vars["todoId"]); err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		panic(err)
	}

	ModelUpdateTodo(todoID, todo)

	_return.Success = true
	_return.Message = "Item added successfully"
	_return.Todos = todos
	res.SendJSON(_return)
}
