package core

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/user/learn-go-myself/controllers"
	"log"
	"net/http"
)

type Server struct {
	router   *mux.Router
	database *sql.DB
}

func Run() *Server {
	s := &Server{
		router: mux.NewRouter(),
	}
	s.RegisterRoutes()

	database, err := GetDB()
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	controllers.SetDB(database)
	log.Fatal(http.ListenAndServe(":3000", s.router))
	return s
}
