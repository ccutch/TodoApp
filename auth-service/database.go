package auth

import "todos/api"

// User Database Logic - TODO replace with SQLite or Postgres
var userDB = map[string]*api.User{}

func lookup(email string) *api.User {
	return userDB[email]
}

func save(u *api.User) error {
	userDB[u.Email] = u
	return nil // Future proof?
}
