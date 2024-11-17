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
	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		userId INTEGER
	)
	`

	_, err := DB.Exec(createEventTable)
	if err != nil {
		panic("Table could not be created! ")
	}
}
