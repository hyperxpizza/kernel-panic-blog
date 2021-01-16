package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/hyperxpizza/kernel-panic-blog/server/models"
)

func CreatePost(title, content, slug string, userID int, subtitle *string) error {
	stmt, err := db.Prepare(`INSERT INTO posts VALUES(DEFAULT, $1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userID, title, subtitle, content, time.Now(), time.Now(), slug)
	if err != nil {
		return err
	}

	return nil

}

func CheckIfPostExists(title string) bool {
	err := db.QueryRow(`SELECT title FROM posts WHERE title=$1`, title).Scan(&title)
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

func GetAllPosts() ([]*models.Post, error) {
	var posts []*models.Post

	rows, err := db.Query(`SELECT * FROM posts;`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var post models.Post

		err = rows.Scan(
			&post.ID,
			&post.AuthorID,
			&post.Tilte,
			&post.Subtitle,
			&post.Content,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.Slug,
		)
		if err != nil {
			return nil, err
		}

		posts = append(posts, &post)
	}

	return posts, nil
}

func GetPostByID(postID int) (*models.Post, error) {
	var post models.Post

	err := db.QueryRow(`SELECT * FROM posts;`).Scan(
		&post.ID,
		&post.AuthorID,
		&post.Tilte,
		&post.Subtitle,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.Slug,
	)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
