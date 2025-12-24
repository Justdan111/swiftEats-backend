package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/Justdan111/swiftEats-backend/internal/db"
	"github.com/Justdan111/swiftEats-backend/internal/middleware"
	"github.com/Justdan111/swiftEats-backend/internal/restaurant"
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

	// Initialize SQLC queries
	queries := db.New(dbStore.DB)

	// Initialize middlewares
	authMiddleware := middleware.AuthMiddleware(jwtSecret)
	adminMiddleware := middleware.AdminMiddleware(jwtSecret)

	// ============ USER MODULE ============
	userRepo := user.NewRepository(dbStore.DB)
	userService := user.NewService(userRepo, jwtSecret)
	userHandler := user.NewHandler(userService)

	// ============ RESTAURANT MODULE ============
	restaurantRepo := restaurant.NewRepository(queries)
	restaurantService := restaurant.NewService(restaurantRepo)
	restaurantHandler := restaurant.NewHandler(restaurantService)

	// ============ ROUTER SETUP ============
	r := mux.NewRouter()

	// Health Check Route
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}).Methods("GET")

	// Public auth endpoints
	r.HandleFunc("/api/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/api/login", userHandler.Login).Methods("POST")

	// 🔒 Protected user route
	r.Handle("/api/me", authMiddleware(http.HandlerFunc(userHandler.Me))).Methods("GET")

	// Restaurant routes (user reads + admin writes)
	restaurant.RegisterRoutes(r, restaurantHandler, adminMiddleware)

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
