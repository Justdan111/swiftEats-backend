package user

import (
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateUser(email, hashedPassword string) error {
		_, err := r.DB.Exec(`INSERT INTO users (email, password_hash) VALUES ($1, $2)`, email, hashedPassword)
		return err
} 

func (r *Repository) GetUserByEmail(email string) (*User, error) {
	row := r.DB.QueryRow(`SELECT id, email, password_hash, created_at FROM users WHERE email=$1`, email)
	u := &User{}
	err := row.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt)
	return u, err
}