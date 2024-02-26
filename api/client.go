package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type Client struct {
	// Host url for api server
	Host string

	// Optional token for user
	Token string
}

// Create a new client to connect to given host, with given identity
func NewClient(host, token string) *Client {
	return &Client{host, token}
}

func (c *Client) CallAPI(method, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, c.Host+path, body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create http request")
	}
	if c.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make http request")
	}
	if res.StatusCode != 200 {
		b, _ := io.ReadAll(res.Body)
		return nil, errors.Errorf("failed http request; response body: %s", b)
	}
	return res, nil
}
