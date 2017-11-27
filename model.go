package main

import (
	"fmt"
	"time"
)

// Todo define structure
type Todo struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos defien Todos
type Todos []Todo

// Define vars
var currentID int
var todos Todos

// Give us some seed data
func init() {
	ModelCreateTodo(Todo{Name: "Make presentation"})
	ModelCreateTodo(Todo{Name: "Host meetup"})
}

// ModelFindTodo find Todo by id
func ModelFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Todo{}
}

// ModelCreateTodo create a Todo item
func ModelCreateTodo(t Todo) Todo {
	currentID++
	t.Id = currentID
	todos = append(todos, t)
	return t
}

// ModelUpdateTodo update an item
func ModelUpdateTodo(id int, todo Todo) Todos {
	for index, t := range todos {
		if t.Id == id {
			// todos[index].Name = todo.Name
			todos[index].Completed = todo.Completed
			return nil
		}
	}
	// return empty Todo if not found
	return todos
}

// ModelDestroyTodo delete an item by id
func ModelDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
