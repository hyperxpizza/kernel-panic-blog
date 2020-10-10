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
	Subtitle  string    `json:"subtitle"`
	Content   string    `json:"content"`
	Slug      string    `json:"slug"`
	Lang      string    `json:"lang"`
	AuthorID  uuid.UUID `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Views     int       `json:"views"`
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
		Views:     1,
	}

	stmt, err := db.Prepare(`INSERT INTO posts VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	_, err = stmt.Exec(post.ID, post.Title, post.Subtitle, post.Content, post.Slug, post.Lang, post.AuthorID, post.CreatedAt, post.UpdatedAt, post.Views)
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
		return false
	default:
		return true
	}

}

func GetAllPosts() ([]Post, error) {
	var posts []Post

	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var post Post
		err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Subtitle,
			&post.Content,
			&post.Slug,
			&post.Lang,
			&post.AuthorID,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.Views,
		)

		if err != nil {
			log.Fatal(err)
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
		return false
	default:
		return true
	}

}

func DeletePost(id uuid.UUID) error {

	stmt, err := db.Prepare(`DELETE FROM posts WHERE post_id = $1`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func GetPostBySlug(slug string) (*Post, error) {
	var post Post

	err := db.QueryRow(`SELECT * FROM posts WHERE slug = $1`, slug).Scan(
		&post.ID,
		&post.Title,
		&post.Subtitle,
		&post.Content,
		&post.Slug,
		&post.Lang,
		&post.AuthorID,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.Views,
	)

	switch {
	case err == sql.ErrNoRows:
		log.Fatal(err)
		return nil, err
	case err != nil:
		log.Fatal(err)
		return nil, err
	default:
		return &post, nil
	}

}
