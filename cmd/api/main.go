package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Justdan111/swiftEats-backend/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
    _ = godotenv.Load() // Load .env file if it exists

    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        dbURL = "postgres://dev:dev@localhost:5432/swiftEats?sslmode=disable"
    }

    st, err := store.NewStore(dbURL)
    if err != nil {
        log.Fatal("DB connect error:", err)
    }
    defer st.Close()

	   // log the database version
    var version string
    err = st.DB.QueryRow(context.Background(), "SELECT version()").Scan(&version)
    if err != nil {
        log.Println("Warning: Could not get DB version:", err)
    } else {
        fmt.Println("Connected to:", version)
    }

    r := chi.NewRouter()
    r.Use(middleware.Logger)

    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(`{"status":"ok"}`))
    })

    fmt.Println("🚀 Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}