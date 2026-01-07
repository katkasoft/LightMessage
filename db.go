package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var Db *sql.DB

func initDB() {
	var err error
	Db, err = sql.Open("sqlite", "lightmessage.db")
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT UNIQUE NOT NULL,
        email TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL
    );`

	_, err = Db.Exec(query)
	if err != nil {
		log.Fatal("Error while creating table:", err)
	}

	log.Println("DB is initialised")
}
