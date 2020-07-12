package core

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"
)
var (
	m  sync.Mutex
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "andon"
)

func (s *server) GetDB() error {
	m.Lock()
	defer m.Unlock()
	var err error

	if s.database == nil {
		connStr := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		s.database, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Print("Error connecting to db", err)
			return fmt.Errorf("Error connecting to db %v ", err)
		}
	}
	err = s.database.Ping()
	if err != nil {
		panic(err)
	}
	return nil
}

