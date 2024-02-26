package auth

import (
	"net/http"
	"os"
	"strings"
	"time"
	"todos/api"

	"github.com/pkg/errors"
)

var enabledCookies = os.Getenv("COOKIES_ENABLED") == "true"

// User Session Logic - attaching cookie for user
func attach(w http.ResponseWriter, u *api.User) string {
	token := token(u)
	if enabledCookies {
		cookie := http.Cookie{
			Name:     cookieName,
			Value:    token,
			Expires:  time.Now().Add(72 * time.Hour),
			Secure:   true,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, &cookie)
	}
	w.Header().Set("Auth-Token", token)
	return token
}

func Current(r *http.Request) (user *api.User, err error) {
	if header := r.Header.Get("Authorization"); strings.HasPrefix(header, "Bearer ") {
		token := strings.Split(header, " ")[1]
		if user, err = authorize(token); err != nil {
			return nil, errors.Wrap(err, "invalid header")
		}
	}
	if cookie, err := r.Cookie(cookieName); enabledCookies && err == nil {
		if user, err = authorize(cookie.Value); err != nil {
			return nil, errors.Wrap(err, "invalid cookie")
		}
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
