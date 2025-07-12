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

// Opens datbase connection
func Connect(database string) *sql.DB {
	return connect(database)
}
