package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitialDB() {
	var err error
	DB, err = sql.Open("sqlite3", "app.db")
	if err != nil {
		panic("data base could not run!")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUserTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
		)`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic("Table could not be created! ")
	}

	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		userId INTEGER,
		FOREIGN KEY(userID) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventTable)
	if err != nil {
		panic("Table could not be created! ")
	}

	createRegTable := `
	CREATE TABLE IF NOT EXISTS erg (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		eventId INTEGER,
		userId INTEGER,
		FOREIGN KEY(userID) REFERENCES users(id)
		FOREIGN KEY(eventID) REFERENCES events(id)
	)
	`

	_, err = DB.Exec(createRegTable)
	if err != nil {
		panic("Table could not be created! ")
	}
}
