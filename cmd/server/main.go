package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/Justdan111/swiftEats-backend/internal/user"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DATABASE_URL")
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := user.NewRepository(db)
	service := user.NewService(repo, jwtSecret)
	handler := user.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/api/register", handler.Register).Methods("POST")
	r.HandleFunc("/api/login", handler.Login).Methods("POST")

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
