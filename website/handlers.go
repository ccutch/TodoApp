package website

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
	"todos/api"
	"todos/evnt-service"
)

func (web *Website) client(w http.ResponseWriter, r *http.Request) *api.Client {
	token, err := r.Cookie("auth-token")
	if token == nil || err != nil {
		renderAuthForm(w, errors.New("Session expired"))
		return nil
	}
	return api.NewClient(web.host, token.Value)
}

func (web *Website) handleLogin(w http.ResponseWriter, r *http.Request) {
	client := api.NewClient(web.host, "")
	email, pass := r.FormValue("email"), r.FormValue("password")
	token, err := client.Signin(email, pass)
	if err != nil {
		renderAuthForm(w, err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "auth-token",
		Value:    token,
		Expires:  time.Now().Add(72 * time.Hour),
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
	})
	renderTodoList(w, api.NewClient(web.host, token), nil)
}

func (web *Website) handleRegister(w http.ResponseWriter, r *http.Request) {
	client := api.NewClient(web.host, "")
	email, pass := r.FormValue("email"), r.FormValue("password")
	token, err := client.Signup(email, pass)
	if err != nil {
		renderAuthForm(w, err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "auth-token",
		Value:    token,
		Expires:  time.Now().Add(72 * time.Hour),
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
	})
	renderTodoList(w, api.NewClient(web.host, token), nil)
}

func (web *Website) handleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth-token",
		Value:    "",
		Expires:  time.Now(),
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
	})
	w.Header().Set("Hx-Refresh", "true")
	w.WriteHeader(http.StatusNoContent)
}

func (web *Website) handleCreate(w http.ResponseWriter, r *http.Request) {
	if client := web.client(w, r); client != nil {
		todo, err := client.CreateTodo(r.FormValue("text"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var buf bytes.Buffer
		renderTodoItem(io.MultiWriter(w, &buf), todo)
		go func(html string) {
			user, _ := client.Identity()
			evnt.Broadcast(user.ID, "todo-created", html)
		}(buf.String())
	}
}

func (web *Website) handleToggle(w http.ResponseWriter, r *http.Request) {
	if client := web.client(w, r); client != nil {
		t, err := client.FetchTodo(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todo, err := client.UpdateTodo(t.ID, t.Text, !t.Complete)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var buf bytes.Buffer
		renderTodoItem(io.MultiWriter(w, &buf), todo)
		go func(html string) {
			user, _ := client.Identity()
			evnt.Broadcast(user.ID, "todo-updated", html)
		}(buf.String())
	}
}

func (web *Website) handleClear(w http.ResponseWriter, r *http.Request) {
	if client := web.client(w, r); client != nil {
		_, err = client.DeleteTodo(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, "")
	}
}
