package database

import (
	"database/sql"
	"fmt"
	"log"

	//Postgresql driver
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host = "localhost"
	port = 5432
)

func InitDB(user, password, dbname string) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	log.Println(psqlInfo)

	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}

	db = database
	log.Println("[+] Connected to the database")

}
