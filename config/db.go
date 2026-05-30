package config

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// DB is our global database connection pool accessible by other packages
var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./storage.db")
	if err != nil {
		panic(err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        age INTEGER
    );`

	_, err = DB.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
}
