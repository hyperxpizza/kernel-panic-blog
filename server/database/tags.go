package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/gosimple/slug"
)

/*
func GetTagsByPostID(id int) ([]models.Tag, error) {

}

func InsertTag(name, slug string, postID int) error {

}
*/

func CheckIfTagExists(tag string) bool {
	err := db.QueryRow(`SELECT tagName from tags WHERE tagName=$1`, tag).Scan(&tag)
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

func CheckIfTagMapExists(tagID, postID int) bool {
	err := db.QueryRow(`SELECT tagID, postID FROM tagamap WHERE tagID=$1, postID=$2`, tagID, postID).Scan(&tagID, postID)
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

func GetIDbyTagName(tag string) (*int, error) {
	var id int
	err := db.QueryRow(`SELECT id FROM tags WHERE tagName=$1`, tag).Scan(id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func InsertTag(tag string) error {
	//create slug
	slug := slug.Make(tag)

	stmt, err := db.Prepare(`INSERT INTO tags VALUES(DEFAULT, $1, $2, $3, $4)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(tag, slug, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil

}
