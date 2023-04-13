// Package database gives access to opening and querying available databases
package database

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// OpenPostgreSQLConnection returns a db connection to a postgres database
func OpenPostgreSQLConnection() (*sqlx.DB, error) {
	maxOpenConns, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	maxIdleConns, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	connMaxLifetime, _ := strconv.Atoi(os.Getenv("DB_CONN_MAX_LIFETIME"))

	db, err := sqlx.Connect("postgres", os.Getenv("DB_CONN_STRING"))
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(time.Duration(connMaxLifetime))

	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error, database ping failed, %w", err)
	}

	return db, nil
}
