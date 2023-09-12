package db

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=elio sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
