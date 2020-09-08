package database

import (
	"database/sql"
	"fmt"
	"log"

	//Postgresql driver
	_ "github.com/lib/pq"
)

var databaseInstance *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "pizza"
	password = "Wojtekfoka1"
	dbname   = "kernelpanicblog"
)

func InitDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	databaseInstance = db
	log.Println("Connected to the database")

}
