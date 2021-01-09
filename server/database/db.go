package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	//Postgresql driver

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {

	//read enviroment variables
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

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
