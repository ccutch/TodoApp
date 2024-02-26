package api

// User model for storing credentials and profile data
type User struct {
	ID       string
	Name     string
	Email    string
	PassHash []byte `json:"-"`
}
