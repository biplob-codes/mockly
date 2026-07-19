package database

import (
	"database/sql"
 
	"fmt"
	_ "modernc.org/sqlite"
)



func Connect(path string) (*sql.DB, error) {
	dbPath:=path+"?_pragma=foreign_keys(1)"
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}
	return db, nil
}
