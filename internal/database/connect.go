package database

import (
	"database/sql"
 
	"fmt"
	_ "modernc.org/sqlite"
)



func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "mockly.db")
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}
	return db, nil
}
