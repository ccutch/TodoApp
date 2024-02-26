package website

import (
	"net/http"
)

type Website struct {
	host string
	mux  *http.ServeMux
}

func Connect(host string) *Website {
	web := Website{host, http.NewServeMux()}
	web.mux.HandleFunc("GET /{$}", web.homepage)
	web.mux.HandleFunc("POST /login", web.handleLogin)
	web.mux.HandleFunc("POST /register", web.handleRegister)
	web.mux.HandleFunc("POST /logout", web.handleLogout)
	web.mux.HandleFunc("POST /create", web.handleCreate)
	web.mux.HandleFunc("POST /toggle/{id}", web.handleToggle)
	web.mux.HandleFunc("POST /clear/{id}", web.handleClear)
	return &web
}

func (web *Website) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	web.mux.ServeHTTP(w, r)
}
