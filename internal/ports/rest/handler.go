package rest

import (
	"encoding/json"
	"net/http"

	"github.com/AspenFresh/lab4-webapp/internal"
)

type Handler struct {
	service *internal.UserService
}

func NewHandler(service *internal.UserService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user internal.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	createdUser, err := h.service.CreateUser(r.Context(), user)
	if err != nil {
		http.Error(w, "failed to create user: "+err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}
