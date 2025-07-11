package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

// Opens database connection
func connect() *sql.DB {
	db, e := sql.Open("sqlite", "./elf.db")
	if e != nil {
		panic(e)
	}
	return db
}
