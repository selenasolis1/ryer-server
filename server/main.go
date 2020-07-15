package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	// ListenPort is the local port that will be used to start the server
	ListenPort int
)

// ExampleServer is a type that will be used to store the route information
// for this specific server type.
type ExampleServer struct {
	Router *http.ServeMux
}

// ServeHTTP is a method that utilizes the writer and reader specified by the
// router type.
func (s *ExampleServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func main() {
	flag.IntVar(&ListenPort, "p", 6500, "Which port to listen on as the server")
	flag.Parse()

	// Server
	server := http.Server{Addr: fmt.Sprintf("0.0.0.0:%d", ListenPort)}

	// Mux instead of a struct with ServeHTTP for the example
	srv := &ExampleServer{
		Router: http.NewServeMux(),
	}
	srv.Routes()
	server.Handler = srv

	log.Printf("Greeting server listening on port %d\n", ListenPort)
	_ = server.ListenAndServe()
}
