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

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "andon"
)

func GetDB() (*sql.DB, error) {
	m.Lock()
	var err error

	if db == nil {
		connStr := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Print("Error connecting to db", err)
			return nil, fmt.Errorf("Error connecting to db %v ", err)
		}
	}
	m.Unlock()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db, nil
}

