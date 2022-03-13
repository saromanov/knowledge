package rest

import "time"

type Page struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	AuthorID  string    `json:"author_id"`
	Links     []string  `json:"links"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
