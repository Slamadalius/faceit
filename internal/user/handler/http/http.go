package http

import (
	"encoding/json"
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

	router.HandleFunc("/user", handler.createUser).Methods(http.MethodPost)
	router.HandleFunc("/user/{id}", handler.updateUser).Methods(http.MethodPut)
	router.HandleFunc("/user/{id}", handler.deleteUser).Methods(http.MethodDelete)
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	user := entity.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		write(w, []byte(`{"error": "bad request"}`), http.StatusBadRequest)

		return
	}

	if err := h.Service.CreateUser(r.Context(), user); err != nil {
		write(w, []byte(`{"error": "internal server error"}`), http.StatusInternalServerError)

		return
	}

	write(w, []byte(`{"success": "user created succesfully"}`), http.StatusOK)
}

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	var (
		user   = entity.User{}
		params = mux.Vars(r)
		userID = params["id"]
	)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		write(w, []byte(`{"error": "bad request"}`), http.StatusBadRequest)

		return
	}

	if err := h.Service.UpdateUser(r.Context(), userID, user); err != nil {
		write(w, []byte(`{"error": "internal server error"}`), http.StatusInternalServerError)

		return
	}

	write(w, []byte(`{"success": "user updated succesfully"}`), http.StatusOK)
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	var (
		params = mux.Vars(r)
		userID = params["id"]
	)

	if err := h.Service.DeleteUser(r.Context(), userID); err != nil {
		write(w, []byte(`{"error": "internal server error"}`), http.StatusInternalServerError)

		return
	}

	write(w, []byte(`{"success": "user deleted succesfully"}`), http.StatusOK)
}
