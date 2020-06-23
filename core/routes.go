package core

import (
	"context"
	"github.com/user/learn-go-myself/controllers"
	"net/http"

)

func (s *Server) RegisterRoutes() {

	ctx := context.Background()
	s.router.HandleFunc("/json", controllers.ReadJSON()).Methods("POST")
	s.router.HandleFunc("/wc", controllers.Get(ctx)).Methods("GET")
	s.router.Handle("/", http.NotFoundHandler())

}