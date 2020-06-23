package core

import (
	"github.com/user/learn-go-myself/controllers"
	"net/http"

)

func (s *Server) RegisterRoutes() {

	s.router.HandleFunc("/json", controllers.ReadJSON()).Methods("POST")
	s.router.HandleFunc("/wc", controllers.Get()).Methods("GET")
	s.router.Handle("/", http.NotFoundHandler())

}