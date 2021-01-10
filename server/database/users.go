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

func GetUsersPassword(username string) string {
	var password string
	err := db.QueryRow(`SELECT passwordHash FROM users WHERE username = $1`, username).Scan(&password)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return password
}

func GetAdminAndID(username string) (bool, int) {
	var isAdmin bool
	var id int
	err := db.QueryRow(`SELECT isAdmin, id FROM users WHERE username=$1`, username).Scan(&isAdmin, &id)
	if err != nil {
		log.Panic(err)
	}

	return isAdmin, id
}
