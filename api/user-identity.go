package api

import (
	"encoding/json"
	"net/http"
)

// Identity of the user for the api client
func (c *Client) Identity() (*User, error) {
	// req, err := http.NewRequest(http.MethodGet, c.Host+"/auth/identity", nil)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to create http request")
	// }
	// if c.Token != "" {
	// 	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	// }
	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to make http request")
	// }
	// if res.StatusCode != 200 {
	// 	b, _ := io.ReadAll(res.Body)
	// 	return nil, errors.Errorf("failed to get identity: %s", b)
	// }
	res, err := c.CallAPI(http.MethodGet, "/auth/identity", nil)
	if err != nil {
		return nil, err
	}
	var user User
	return &user, json.NewDecoder(res.Body).Decode(&user)
}
