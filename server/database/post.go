package database

import (
	"log"
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

func InsertPost(title, subtitle, content, slug string, authorID uuid.UUID) (*Post, error) {
	post := Post{
		ID:        uuid.Must(uuid.NewV1()),
		Title:     title,
		Subtitle:  subtitle,
		Content:   content,
		Slug:      slug,
		AuthorID:  authorID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	stmt, err := db.Prepare(`INSERT INTO posts VALUES($1, $2, $3, $4, $5, $6, $7, $8)`)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	_, err = stmt.Exec(post.CreatedAt, post.UpdatedAt, post.ID, post.Title, post.Subtitle, post.Content, post.Slug, post.AuthorID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &post, nil
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
