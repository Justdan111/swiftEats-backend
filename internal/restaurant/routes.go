package restaurant

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, handler *Handler) {
	r.HandleFunc("/api/restaurants", handler.ListRestaurants).Methods("GET")
	r.HandleFunc("/api/restaurants/{id}", handler.GetRestaurant).Methods("GET")
	r.HandleFunc("/api/restaurants/{id}/menu", handler.GetMenu).Methods("GET")
}
