package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// DBManager wraps a sql.DB object for interacting with the database.
type DBManager struct {
	db *sql.DB
}

// NewDBManager initializes and returns a new DBManager.
func NewDBManager(dataSourceName string) (*DBManager, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return &DBManager{db}, nil
}

// Query executes a SQL query with the provided arguments and returns the result rows.
func (manager *DBManager) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return manager.db.Query(query, args...)
}

// QueryRow executes a SQL query with the provided arguments and returns a single row.
func (manager *DBManager) QueryRow(query string, args ...interface{}) *sql.Row {
	return manager.db.QueryRow(query, args...)
}

// Exec executes a SQL query with the provided arguments and returns the result.
func (manager *DBManager) Exec(query string, args ...interface{}) (sql.Result, error) {
	return manager.db.Exec(query, args...)
}

// Close closes the database connection.
func (manager *DBManager) Close() error {
	return manager.db.Close()
}
