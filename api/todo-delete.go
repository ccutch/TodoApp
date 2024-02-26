package api

import (
	"encoding/json"
	"net/http"
)

func (c *Client) DeleteTodo(id string) (*Todo, error) {
	res, err := c.CallAPI(http.MethodDelete, "/todo/"+id, nil)
	if err != nil {
		return nil, err
	}
	var todo Todo
	return &todo, json.NewDecoder(res.Body).Decode(&todo)
}
