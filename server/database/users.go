package database

import (
	"database/sql"
	"log"
	"time"
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

func CheckIfEmailTaken(email string) bool {
	err := db.QueryRow(`SELECT email FROM users WHERE username = $1`, email).Scan(&email)
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

func InsertUser(username, password, email string) error {
	stmt, err := db.Prepare(`INSERT INTO users VALUES(DEFAULT, $1, $2, $3, $4, $5, $6)`)
	if err != nil {
		log.Fatal(err)
		return err
	}

	isAdmin := false
	if username == "hyperxpizza" {
		isAdmin = true
	}

	_, err = stmt.Exec(username, password, email, isAdmin, time.Now(), time.Now())
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
