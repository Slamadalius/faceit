package http

import (
	"log"
	"net/http"
)

func write(w http.ResponseWriter, data []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if _, err := w.Write(data); err != nil {
		log.Println("ERR: ", err)
	}
}
