package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(handler *Handler, authMiddleware func(http.Handler) http.Handler) http.Handler {
	r := chi.NewRouter()

	r.Post("/register", handler.Register)
	r.Post("/login", handler.Login)

	// Protected routes
	r.Group(func(pr chi.Router) {
		pr.Use(authMiddleware)

		pr.Get("/me", handler.Me)
	})

	return r
}