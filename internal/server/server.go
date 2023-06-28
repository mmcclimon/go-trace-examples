package server

import (
	"context"
	"log"
	"net/http"
)

type Server struct {
	server http.Server
}

type responderFunc = func(http.ResponseWriter, *http.Request)

func ServeHTTP(addr string, responder responderFunc) *Server {
	http.HandleFunc("/", responder)

	s := &Server{http.Server{Addr: addr}}
	go s.serveHTTP()
	return s
}

func (s *Server) Shutdown() {
	err := s.server.Shutdown(context.Background())
	if err != nil {
		log.Print(err)
	}
}

func (s *Server) serveHTTP() {
	err := s.server.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
