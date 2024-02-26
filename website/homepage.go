package website

import (
	"net/http"
	"todos/api"
)

func (web *Website) homepage(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Token  string
		Error  string
		Todos  []api.Todo
		UserID string
	}
	if cookie, err := r.Cookie("auth-token"); cookie != nil && err == nil {
		client := api.NewClient(web.host, cookie.Value)
		if body.Todos, err = client.ListTodos(); err != nil {
			body.Error = "Session expired"
		} else {
			body.Token = cookie.Value
		}
		if body.Error == "" {
			user, _ := client.Identity()
			body.UserID = user.ID
		}
	}
	render(w, "homepage.html", body)
}
