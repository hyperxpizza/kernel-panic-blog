package database

import (
	"time"

	"github.com/gofrs/uuid"
)

type Post struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	Content   string    `json:"content"`
	Slug      string    `json:"slug"`
	AuthorID  uuid.UUID `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func InsertPost(title, subtitle, content, slug string, authorID uuid.UUID) {

}

func GetAllPosts() ([]Post, error) {
	var posts []Post

	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var post Post
		err = rows.Scan(&post.CreatedAt, &post.UpdatedAt, &post.ID, &post.Title, &post.Subtitle, &post.Content, &post.Slug, &post.AuthorID)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil

}
