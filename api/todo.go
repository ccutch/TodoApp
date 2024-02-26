package api

import "time"

type Todo struct {
	ID        string
	UserID    string
	Text      string
	Complete  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
