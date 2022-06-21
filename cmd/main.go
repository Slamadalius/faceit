package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/faceit", handler).Methods("GET")

	server := &http.Server{
		Handler: http.TimeoutHandler(router, 15*time.Second, "response timeout"),
		Addr:    ":8080",
	}

	go func() {
		log.Print("server listening on port :8080")
		if err := server.ListenAndServe(); err != nil {
			log.Print("server closed")
		} else {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Print("server shutting down")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
