package auth

import (
	"encoding/json"
	"net/http"
	"todos/api"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Name     string
	Email    string
	Password string
}

func parse(r *http.Request) (c *Credentials, err error) {
	err = json.NewDecoder(r.Body).Decode(&c)
	return c, errors.Wrap(err, "Failed to parse credentials")
}

func (c *Credentials) create() (*api.User, error) {
	if _, ok := userDB[c.Email]; ok {
		return nil, errors.New("account already exists for this email")
	}
	passHash, err := bcrypt.GenerateFromPassword([]byte(c.Password), 0)
	if err != nil {
		return nil, err
	}
	user := api.User{
		ID:       api.GenerateID(11),
		Name:     c.Name,
		Email:    c.Email,
		PassHash: passHash,
	}
	return &user, save(&user)
}

func (c *Credentials) authenticate() (*api.User, error) {
	user := lookup(c.Email)
	if user == nil {
		return nil, errors.New("no account exists for this email")
	}
	return user, bcrypt.CompareHashAndPassword(user.PassHash, []byte(c.Password))
}
