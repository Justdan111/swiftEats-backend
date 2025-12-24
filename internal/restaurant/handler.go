package restaurant

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) ListRestaurants(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.ListRestaurants(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func (h *Handler) GetRestaurant(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	data, err := h.service.GetRestaurant(r.Context(), id)
	if err != nil {
		http.Error(w, "restaurant not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func (h *Handler) GetMenu(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	menu, err := h.service.GetMenu(r.Context(), id)
	if err != nil {
		http.Error(w, "menu not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(menu)
}
