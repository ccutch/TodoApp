package rest

import (
	"encoding/json"
	"net/http"
	"todos/api"
	"todos/auth-service"
)

func directoryTodos(w http.ResponseWriter, r *http.Request) {
	user, err := auth.Current(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	var myTodos []*api.Todo
	for _, t := range todoDB {
		if t.UserID == user.ID {
			myTodos = append(myTodos, t)
		}
	}
	json.NewEncoder(w).Encode(myTodos)
}

func lookupTodo(w http.ResponseWriter, r *http.Request) {
	user, err := auth.Current(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	todo, ok := todoDB[r.PathValue("id")]
	if !ok {
		http.Error(w, "todo not found", http.StatusNotFound)
		return
	}
	if todo.UserID != user.ID {
		http.Error(w, "who are you?", http.StatusForbidden)
		return
	}
	json.NewEncoder(w).Encode(todo)
}
