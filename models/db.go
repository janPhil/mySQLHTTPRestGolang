package models

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

// NewDB builds the connection to the database and returns a handle
func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:test@tcp(127.0.0.1:3306)/employees")
	if err != nil {
		fmt.Printf("Could not connect to database %v", err)
	}
	return db
}
