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



// ListRestaurants returns all restaurants
func (h *Handler) ListRestaurants(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.ListRestaurants(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch restaurants", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// GetRestaurant returns a single restaurant
func (h *Handler) GetRestaurant(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	data, err := h.service.GetRestaurant(r.Context(), id)
	if err != nil {
		http.Error(w, "restaurant not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// GetMenu returns menu items for a restaurant
func (h *Handler) GetMenu(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	menu, err := h.service.GetMenuByRestaurantID(r.Context(), id)
	if err != nil {
		http.Error(w, "menu not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menu)
}



// CreateRestaurant creates a new restaurant (admin only)
func (h *Handler) CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	var input RestaurantInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input
	if input.Name == "" || input.Address == "" {
		http.Error(w, "name and address are required", http.StatusBadRequest)
		return
	}

	restaurant, err := h.service.CreateRestaurant(r.Context(), input)
	if err != nil {
		http.Error(w, "failed to create restaurant", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(restaurant)
}

// UpdateRestaurant updates a restaurant (admin only)
func (h *Handler) UpdateRestaurant(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var input RestaurantInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input
	if input.Name == "" || input.Address == "" {
		http.Error(w, "name and address are required", http.StatusBadRequest)
		return
	}

	restaurant, err := h.service.UpdateRestaurant(r.Context(), id, input)
	if err != nil {
		http.Error(w, "failed to update restaurant", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurant)
}

// DeleteRestaurant deletes a restaurant (admin only)
func (h *Handler) DeleteRestaurant(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := h.service.DeleteRestaurant(r.Context(), id)
	if err != nil {
		http.Error(w, "failed to delete restaurant", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}



// CreateMenuItem creates a new menu item (admin only)
func (h *Handler) CreateMenuItem(w http.ResponseWriter, r *http.Request) {
	var input MenuItemInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input
	if input.Name == "" || input.PriceCents < 0 || input.RestaurantID == nil {
		http.Error(w, "name, price_cents, and restaurant_id are required", http.StatusBadRequest)
		return
	}

	item, err := h.service.CreateMenuItem(r.Context(), input)
	if err != nil {
		http.Error(w, "failed to create menu item", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// UpdateMenuItem updates a menu item (admin only)
func (h *Handler) UpdateMenuItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var input MenuItemInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input
	if input.Name == "" || input.PriceCents < 0 {
		http.Error(w, "name and price_cents are required", http.StatusBadRequest)
		return
	}

	item, err := h.service.UpdateMenuItem(r.Context(), id, input)
	if err != nil {
		http.Error(w, "failed to update menu item", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// DeleteMenuItem deletes a menu item (admin only)
func (h *Handler) DeleteMenuItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := h.service.DeleteMenuItem(r.Context(), id)
	if err != nil {
		http.Error(w, "failed to delete menu item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// UpdateMenuItemAvailability toggles menu item availability (admin only)
func (h *Handler) UpdateMenuItemAvailability(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var input struct {
		IsAvailable bool `json:"is_available"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := h.service.UpdateMenuItemAvailability(r.Context(), id, input.IsAvailable)
	if err != nil {
		http.Error(w, "failed to update menu item availability", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"is_available": input.IsAvailable})
}
