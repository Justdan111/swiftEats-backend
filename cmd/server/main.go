package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/Justdan111/swiftEats-backend/internal/middleware"
	"github.com/Justdan111/swiftEats-backend/internal/store"
	"github.com/Justdan111/swiftEats-backend/internal/user"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DATABASE_URL")
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	// Initialize database with sql.DB using pgx driver
	dbStore, err := store.NewStore(dbURL)
	if err != nil {
		log.Fatal("❌ Cannot connect to database:", err)
	}
	defer dbStore.Close()
	fmt.Println("🚀 Connected to database successfully")

	// Initialize user module with database
	repo := user.NewRepository(dbStore.DB)
	service := user.NewService(repo, jwtSecret)
	handler := user.NewHandler(service)

	authMiddleware := middleware.AuthMiddleware(jwtSecret)

	r := mux.NewRouter()

	// Health Check Route
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}).Methods("GET")

	// Public auth endpoints
	r.HandleFunc("/api/register", handler.Register).Methods("POST")
	r.HandleFunc("/api/login", handler.Login).Methods("POST")

	// 🔒 Protected route
	r.Handle("/api/me", authMiddleware(http.HandlerFunc(handler.Me))).Methods("GET")

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
