package database

import (
	"database/sql"
	"log"
	"time"
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
