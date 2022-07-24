package storage

import "time"

// Page defines page for article
type Page struct {
	ID        int64     `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	AuthorID  int64     `db:"author_id"`
	Links     []Link    `db:"links"`
}

// Author defines author of the article
type Author struct {
	ID        int64     `db:"id"`
	Name      string    `db:"author"`
	CreatedAt time.Time `db:"created_at"`
}

// Link defines urls to external resources
type Link struct {
	ID  int64  `db:"id"`
	URL string `db:"url"`
}
