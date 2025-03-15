package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/DenisGoldiner/webapp/internal"
)

type TravellerHandler struct {
	service internal.Travellers
}

func NewTravellerHandler(service internal.Travellers) TravellerHandler {
	return TravellerHandler{
		service: service,
	}
}

func (h TravellerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetTraveller(w, r)
	case http.MethodPost:
		h.CreateTraveller(w, r)
	case http.MethodDelete:
		h.DeleteTraveller(w, r)
	default:
		msg := fmt.Sprintf("method %s is not supported", r.Method)
		log.Println(msg)
		http.Error(w, msg, http.StatusMethodNotAllowed)
	}
}

func (h TravellerHandler) GetTraveller(w http.ResponseWriter, r *http.Request) {
	log.Println("GetTraveller")

	ctx := r.Context()

	idParam := r.URL.Query().Get("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, fmt.Sprintf("id must be a valid uuid"), http.StatusBadRequest)
		return
	}

	res, err := h.service.GetTraveller(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respTraveller := Traveller{
		ID:        res.ID,
		FirstName: res.FirstName,
		LastName:  res.LastName,
	}

	if err = json.NewEncoder(w).Encode(respTraveller); err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h TravellerHandler) CreateTraveller(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateTraveller")

	if r.Body == nil {
		http.Error(w, "body must not be nil", http.StatusBadRequest)
		return
	}

	var payload CreateTravellerPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, fmt.Sprintf("failed to parse the body %v", err), http.StatusBadRequest)
		return
	}

	log.Println("CreateTravellerPayload", payload)

	h.service.CreateTraveller()

	w.WriteHeader(http.StatusOK)
}

func (h TravellerHandler) DeleteTraveller(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteTraveller")

	h.service.DeleteTraveller()

	w.WriteHeader(http.StatusOK)
}
