package models

import "time"

type Post struct {
	ID        int       `json:"id"`
	AuthorID  int       `json:"author_id"`
	Tilte     string    `json:"title"`
	Subtitle  *string   `json:"subtitle"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Slug      string    `json:"slug"`
}
