package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Init a new router
	r := mux.NewRouter()

	// Define a subrouter for all /api routes
	api := r.PathPrefix("/api").Subrouter()

	// Define api routes
	api.
		Methods("GET").
		Path("/todos").
		HandlerFunc(TodoIndex)

	api.
		Methods("POST").
		Path("/todos").
		HandlerFunc(TodoCreate)

	api.
		Methods("GET").
		Path("/todos/{todoId}").
		HandlerFunc(TodoShow)

	api.
		Methods("DELETE").
		Path("/todos/{todoId}").
		HandlerFunc(TodoDelete)

	api.
		Methods("PUT").
		Path("/todos/{todoId}").
		HandlerFunc(TodoUpdate)

	// Serve all css, js from /assets path
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./client/dist"))))

	// Point all the other requests to the client
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "./client/dist/index.html")
	})

	// Start out server at :8080
	http.ListenAndServe(":8080", handlers.RecoveryHandler()(r))
}