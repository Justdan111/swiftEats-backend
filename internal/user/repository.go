package user

import (
	"context"
	"database/sql"
	"errors"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(ctx context.Context, email, hashedPassword string) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO users (email, password_hash) VALUES ($1, $2)`, email, hashedPassword)
	return err
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	row := r.db.QueryRowContext(ctx, `SELECT id, email, password_hash, created_at FROM users WHERE email=$1`, email)
	u := &User{}
	err := row.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return u, nil
}

func (r *Repository) GetUserByID(ctx context.Context, id string) (*User, error) {
	row := r.db.QueryRowContext(ctx, `SELECT id, email, password_hash, created_at FROM users WHERE id=$1`, id)
	u := &User{}
	err := row.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return u, nil
}
