package core

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/user/learn-go-myself/controllers"
	"log"
	"net/http"
)

type server struct {
	router   *mux.Router
	database *sql.DB
}


func Server() {

	viper.SetConfigFile("./config/config.json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	//port := viper.Get("prod.port") // returns string
	////port := viper.GetInt("prod.port") // returns integer
	//fmt.Printf("Value: %v, Type: %T\n", port, port)

	//var config1 map[string]interface{}
	data := viper.Sub("db")
	fmt.Println(data)

	s := &server{
		router: mux.NewRouter(),
	}
	s.RegisterRoutes()
	err := s.GetDB()
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	controllers.SetDB(s.database)
	defer s.database.Close()
	log.Fatal(http.ListenAndServe(":3000", s.router))
}
