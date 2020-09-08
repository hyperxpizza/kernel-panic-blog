package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

func CheckIfUsernameExists(username string) bool {

	err := databaseInstance.QueryRow(`SELECT username FROM users WHERE username = $1`, username).Scan(&username)
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
	err := databaseInstance.QueryRow("SELECT hashedPassword FROM users WHERE username = $1", username).Scan(&password)
	switch {
	case err == sql.ErrNoRows:
		log.Fatal(err)
	case err != nil:
		log.Fatal(err)
	default:
		return password
	}

	return ""
}

func GetUsersID(username string) uuid.UUID {
	var id uuid.UUID
	_ = databaseInstance.QueryRow(`SELECT id FROM users WHERE username = $1`, username).Scan(&id)

	return id
}

func CheckIfEmailTaken(email string) bool {
	err := databaseInstance.QueryRow(`SELECT email FROM users WHERE email = $1`, email).Scan(&email)
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

func InsertUser(username, password, email string) error {
	stmt, err := databaseInstance.Prepare(`INSERT INTO users VALUES($1, $2, $3, $4, $5, $6)`)
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = stmt.Exec(time.Now(), time.Now(), uuid.Must(uuid.NewV1()), username, password, email)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
