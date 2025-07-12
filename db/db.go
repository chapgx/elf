package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

// Opens database connection
func connect(database string) *sql.DB {
	db, e := sql.Open("sqlite", database)
	if e != nil {
		panic(e)
	}
	return db
}

// Connect Opens datbase connection
func Connect(database string) *sql.DB {
	return connect(database)
}

// Init initializes the database
func Init(database string) error {
	c := connect(database)
	defer c.Close()

	_, e := c.Exec("PRAGMA journal_mode = WAL;")
	if e != nil {
		return e
	}

	_, e = c.Exec(`
	create table if not exists admins (
		id INTEGER PRIMARY KEY,
		uname TEXT,
		masterkey TEXT
	)
	`)

	return e
}
