package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

// Signin an existing user via api client
func (c *Client) Signin(email, password string) (string, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(struct {
		Email, Password string
	}{email, password}); err != nil {
		return "", errors.Wrap(err, "failed to make signin request")
	}
	res, err := http.Post(c.Host+"/auth/signin", "application/json", &buf)
	if err != nil {
		return "", errors.Wrap(err, "failed to make signin request")
	}
	if res.StatusCode != 200 {
		b, _ := io.ReadAll(res.Body)
		return "", errors.Errorf("failed to make signin request: %s", b)
	}
	return res.Header.Get("Auth-Token"), nil
}
