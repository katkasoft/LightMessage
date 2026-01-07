package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

<<<<<<< HEAD
var Db *sql.DB

func initDB() {
	var err error
	Db, err = sql.Open("sqlite", "lightmessage.db")
=======
var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite", "lightmessage.db")
>>>>>>> eb5614cb31287931d6ff6b7d9e1f41f0fdbfcb3d
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

<<<<<<< HEAD
	_, err = Db.Exec(query)
=======
	_, err = db.Exec(query)
>>>>>>> eb5614cb31287931d6ff6b7d9e1f41f0fdbfcb3d
	if err != nil {
		log.Fatal("Error while creating table:", err)
	}

	log.Println("DB is initialised")
}
