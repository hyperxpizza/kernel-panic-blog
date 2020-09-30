package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

type Post struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"string"`
	Content   string    `json:"content"`
	Slug      string    `json:"slug"`
	Lang      string    `json:"lang"`
	AuthorID  uuid.UUID `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func InsertPost(title, subtitle, content, lang, slug string, id uuid.UUID) (*Post, error) {

	post := Post{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ID:        uuid.Must(uuid.NewV1()),
		Title:     title,
		Subtitle:  subtitle,
		Content:   content,
		Slug:      slug,
		AuthorID:  id,
	}

	stmt, err := db.Prepare(`INSERT INTO posts VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	_, err = stmt.Exec(post.ID, post.Title, post.Subtitle, post.Content, post.Slug, post.Lang, post.AuthorID, post.CreatedAt, post.UpdatedAt)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &post, nil
}

func CheckIfPostExists(id uuid.UUID) bool {
	err := db.QueryRow(`SELECT post_id FROM posts WHERE post_id = $1`, id).Scan(&id)
	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		log.Fatal(err)
	default:
		return true
	}

	return true
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
		err = rows.Scan()
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil

}

func CheckIfSlugExists(slug string) bool {
	err := db.QueryRow(`SELECT slug FROM posts WHERE slug = $1`, slug).Scan(&slug)
	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		log.Fatal(err)
	default:
		return true
	}

	return true
}
