package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func (c *Client) CreateTodo(text string) (*Todo, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(struct{ Text string }{text}); err != nil {
		return nil, errors.Wrap(err, "failed to create todo request")
	}
	res, err := c.CallAPI(http.MethodPost, "/todo/", &buf)
	if err != nil {
		return nil, err
	}
	var todo Todo
	return &todo, json.NewDecoder(res.Body).Decode(&todo)
}
