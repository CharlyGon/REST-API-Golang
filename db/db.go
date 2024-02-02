package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to the database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT not null,
		description TEXT,
		location TEXT,
		date_time DATETIME,
		user_id INTEGER
	);
	`
	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}
}
