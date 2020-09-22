package database

import (
	"database/sql"

	"github.com/hyperxpizza/kernel-panic-blog/server/graph/model"
	"vitess.io/vitess/go/vt/log"
)

func GetAllPosts() ([]model.Post, error) {
	var posts []model.Post
	return posts, nil
}

func InsertPost() {

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
