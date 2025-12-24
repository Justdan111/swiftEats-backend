package restaurant

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, handler *Handler, adminAuth func(http.Handler) http.Handler) {
	// ============ USER ROUTES (Public - Read Only) ============
	r.HandleFunc("/api/restaurants", handler.ListRestaurants).Methods("GET")
	r.HandleFunc("/api/restaurants/{id}", handler.GetRestaurant).Methods("GET")
	r.HandleFunc("/api/restaurants/{id}/menu", handler.GetMenu).Methods("GET")

	// ============ ADMIN ROUTES (Protected - Create, Update, Delete) ============
	// Admin restaurant management
	r.Handle("/api/admin/restaurants", adminAuth(http.HandlerFunc(handler.CreateRestaurant))).Methods("POST")
	r.Handle("/api/admin/restaurants/{id}", adminAuth(http.HandlerFunc(handler.UpdateRestaurant))).Methods("PUT")
	r.Handle("/api/admin/restaurants/{id}", adminAuth(http.HandlerFunc(handler.DeleteRestaurant))).Methods("DELETE")

	// Admin menu item management
	r.Handle("/api/admin/menu-items", adminAuth(http.HandlerFunc(handler.CreateMenuItem))).Methods("POST")
	r.Handle("/api/admin/menu-items/{id}", adminAuth(http.HandlerFunc(handler.UpdateMenuItem))).Methods("PUT")
	r.Handle("/api/admin/menu-items/{id}", adminAuth(http.HandlerFunc(handler.DeleteMenuItem))).Methods("DELETE")
	r.Handle("/api/admin/menu-items/{id}/availability", adminAuth(http.HandlerFunc(handler.UpdateMenuItemAvailability))).Methods("PATCH")
}
