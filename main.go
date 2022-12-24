package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var dbClient *sql.DB

const (
	db_host     = "localhost"
	db_port     = 5432
	db_user     = "keluargadb"
	db_password = "keluargadb"
	db_dbname   = "keluargadb"
	http_port   = 8089
	tcp_host    = "localhost"
	tcp_port    = 8088
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		db_host, db_port, db_user, db_password, db_dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	dbClient = db

	go controller()
	go tcpListener()

}
