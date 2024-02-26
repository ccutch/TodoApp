package api

import (
	"encoding/json"
	"net/http"
	"sort"
)

func (c *Client) ListTodos() ([]Todo, error) {
	res, err := c.CallAPI(http.MethodGet, "/todo/", nil)
	if err != nil {
		return nil, err
	}
	var todos []Todo
	if err := json.NewDecoder(res.Body).Decode(&todos); err != nil {
		return nil, err
	}
	sort.Slice(todos, func(i, j int) bool {
		return todos[i].CreatedAt.After(todos[j].CreatedAt)
	})
	return todos, nil
}
