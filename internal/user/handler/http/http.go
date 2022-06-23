package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Slamadalius/faceit/internal/entity"
	"github.com/gorilla/mux"
)

type Handler struct {
	Service entity.UserService
}

func NewUserHandler(router *mux.Router, userService entity.UserService) {
	handler := &Handler{
		Service: userService,
	}

	fmt.Println("New USer")

	router.HandleFunc("/user", handler.createUser).Methods(http.MethodPost)
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	user := entity.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "bad request"}`))

		return
	}

	if err := h.Service.CreateUser(r.Context(), user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "internal error"}`))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"success": "user created succesfully"}`))
}