package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

type Post struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"string"`
	Content   string    `json:"content"`
	Lang      string    `json:"lang"`
	Author    User      `json:"author"`
}

func InsertPost() {
	log.Println("not implemented yet")
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
