package rest

import (
	"net/http"
	"todos/api"
)

func Service() *http.ServeMux {
	mux := http.NewServeMux()
	// Create
	mux.HandleFunc("POST /", createTodo)
	// Retrieve
	mux.HandleFunc("GET /", directoryTodos)
	mux.HandleFunc("GET /{id}", lookupTodo)
	// Update
	mux.HandleFunc("PUT /{id}", updateTodo)
	// Delete
	mux.HandleFunc("DELETE /{id}", deleteTodo)
	return mux
}

// Storing data in memory for now
var todoDB = map[string]*api.Todo{}
