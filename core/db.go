package core

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

var (
	database *sql.DB
)

var (
	db *sql.DB
	m  sync.Mutex
)

func GetDB() (*sql.DB, error) {
	m.Lock()
	var err error

	if db == nil {
		connStr := "user=postgres password=postgres dbname=andon host=localhost sslmode=disable"
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Print("Error connecting to db", err)
			return nil, fmt.Errorf("Error connecting to db %v ", err)
		}
	}
	m.Unlock()
	return db, nil
}

