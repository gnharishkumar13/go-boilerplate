package core

import (
	"context"
	"github.com/user/learn-go-myself/controllers"
	"net/http"
	"os/exec"
)


//func handle(rw http.ResponseWriter, r *http.Request) {
//	rid := r.Header.Get("request-id")
//	if rid != "" {
//		fmt.Fprintf(rw, "Request ID: %s\n", rid)
//	}
//
//	fmt.Fprintln(rw, "Hello, DDD East Anglia")
//}

func addReqID(h http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		out, err := exec.Command("uuidgen").Output()
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		r.Header.Add("request-id", string(out))

		h(rw, r)
	}
}

func (s *server) RegisterRoutes() {

	ctx := context.Background()
	s.router.HandleFunc("/json", addReqID(controllers.ReadJSON())).Methods("POST")
	s.router.HandleFunc("/wc", addReqID(controllers.Get(ctx))).Methods("GET")
	s.router.HandleFunc("/sayhello", controllers.NewName().SayHello()).Methods("GET")
	s.router.Handle("/", http.NotFoundHandler())

}