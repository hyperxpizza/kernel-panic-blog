package database

import (
	"database/sql"
	"log"
)

func CheckIfUsernameExists(username string) bool {
	err := db.QueryRow(`SELECT username FROM users WHERE username = $1`, username).Scan(&username)
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
