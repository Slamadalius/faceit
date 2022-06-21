package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func (s *Server) Start(router http.Handler) {
	s.server = &http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	go func() {
		log.Print("server listening on port :8080")
		if err := s.server.ListenAndServe(); err != nil {
			log.Print("server closed")
		} else {
			log.Fatal(err)
		}
	}()
}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	log.Print("server shutting down")
	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
