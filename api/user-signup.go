package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

// Signup an new user via api client
func (c *Client) Signup(email, password string) (string, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(struct {
		Email, Password string
	}{email, password}); err != nil {
		return "", errors.Wrap(err, "failed to make signup request")
	}
	res, err := http.Post(c.Host+"/auth/signup", "application/json", &buf)
	if err != nil {
		return "", errors.Wrap(err, "failed to make signup request")
	}
	if res.StatusCode != 200 {
		b, _ := io.ReadAll(res.Body)
		return "", errors.Errorf("failed to make signup request: %s", b)
	}
	return res.Header.Get("Auth-Token"), nil
}
