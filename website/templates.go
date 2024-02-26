package website

import (
	"embed"
	"html/template"
	"io"
	"net/http"
	"todos/api"
)

//go:embed templates/* templates/partials/*
var templates embed.FS

var (
	t, err = template.ParseFS(templates, "templates/*.html", "templates/partials/*.html")
	views  = template.Must(t, err)
)

func render(w io.Writer, name string, data interface{}) {
	if err := views.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w.(http.ResponseWriter), err.Error(), http.StatusInternalServerError)
		return
	}
}

func renderAuthForm(w http.ResponseWriter, err error) {
	render(w, "partials/auth-form", struct{ Error string }{err.Error()})
}

func renderTodoList(w io.Writer, c *api.Client, err error) {
	var body struct {
		Todos  []api.Todo
		Error  string
		UserID string
	}
	if err != nil {
		body.Error = err.Error()
	} else {
		if body.Todos, err = c.ListTodos(); err != nil {
			body.Error = "Failed to create todo:" + err.Error()
		} else {
			user, _ := c.Identity()
			body.UserID = user.ID
		}
	}
	render(w, "partials/todo-list", body)
}

func renderTodoItem(w io.Writer, todo *api.Todo) {
	render(w, "partials/todo-item", todo)
}
