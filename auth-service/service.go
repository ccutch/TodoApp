package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"todos/api"
)

var (
	jwtSecret  = os.Getenv("JWT_SECRET")
	cookieName = os.Getenv("COOKIE_NAME")
)

func init() {
	if jwtSecret == "" {
		jwtSecret = api.GenerateID(12)
	}
	if cookieName == "" {
		cookieName = "auth-token"
	}
}

func Service() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /identity", identity)
	mux.HandleFunc("POST /signup", signup)
	mux.HandleFunc("POST /signin", signin)
	return mux
}
func signup(w http.ResponseWriter, r *http.Request) {
	c, err := parse(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := c.create()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	attach(w, user)
	json.NewEncoder(w).Encode(&user)
}

func signin(w http.ResponseWriter, r *http.Request) {
	c, err := parse(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := c.authenticate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	attach(w, user)
	json.NewEncoder(w).Encode(&user)
}

func identity(w http.ResponseWriter, r *http.Request) {
	user, err := Current(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(&user)
}
