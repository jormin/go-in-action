package main

import (
	"database/sql"
	_ "github.com/jormin/go-in-action/chapter3/dbdriver/postgres"
	"gitlab.wcxst.com/jormin/go-tools/log"
)

func main() {
	db, err := sql.Open("postgres", "mydb")
	if err != nil {
		log.Fatal("Open postgres database error: %+v", err)
	}
	log.Info("DB: %+v", db)
	conn, err := db.Driver().Open("localhost:5432")
	if err != nil {
		log.Fatal("Connect to postgres error: %+v", err)
	}
	log.Info("Conn: %+v", conn)
}
