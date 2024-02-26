package rest

import (
	"encoding/json"
	"net/http"
	"todos/auth-service"
)

func updateTodo(w http.ResponseWriter, r *http.Request) {
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
	var body struct {
		Text     string
		Complete *bool
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if body.Text != "" {
		todo.Text = body.Text
	}
	if body.Complete != nil {
		todo.Complete = *body.Complete
	}
	todoDB[todo.ID] = todo
	json.NewEncoder(w).Encode(todo)
}
