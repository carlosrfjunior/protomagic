package database

import "database/sql"

type DBInterface interface {
	Open() (*sql.DB, error)
	// Close() error
	// Query(query string, args ...interface{}) (*sql.Rows, error)
	// Exec(query string, args ...interface{}) (sql.Result, error)
}

type DBCommands interface {
	Run() error
}
