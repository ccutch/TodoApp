package rest

import (
	"encoding/json"
	"net/http"
	"time"
	"todos/api"
	"todos/auth-service"

	"github.com/pkg/errors"
)

func createTodo(w http.ResponseWriter, r *http.Request) {
	user, err := auth.Current(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	var body struct {
		Text string
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		err = errors.Wrap(err, "failed to parse body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo := &api.Todo{
		ID:        api.GenerateID(13),
		UserID:    user.ID,
		Text:      body.Text,
		Complete:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	todoDB[todo.ID] = todo
	json.NewEncoder(w).Encode(todo)
}
