package rest

import (
	"encoding/json"
	"net/http"
	"todos/auth-service"
)

func deleteTodo(w http.ResponseWriter, r *http.Request) {
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
	delete(todoDB, todo.ID)
	json.NewEncoder(w).Encode(todo)
}
