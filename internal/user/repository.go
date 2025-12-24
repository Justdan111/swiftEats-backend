package user

import (
	"context"
	"errors"

	"github.com/Justdan111/swiftEats-backend/internal/db"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	dbTx db.DBTX
}

func NewRepository(dbTx db.DBTX) *Repository {
	return &Repository{
		dbTx: dbTx,
	}
}

func (r *Repository) CreateUser(ctx context.Context, email, hashedPassword string) error {
	// Using DBTX interface methods directly
	_, err := r.dbTx.ExecContext(ctx, `INSERT INTO users (email, password_hash) VALUES ($1, $2)`, email, hashedPassword)
	return err
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	// Using DBTX interface methods directly
	row := r.dbTx.QueryRowContext(ctx, `SELECT id, email, password_hash, created_at FROM users WHERE email=$1`, email)
	u := &User{}
	err := row.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return u, nil
}

func (r *Repository) GetUserByID(ctx context.Context, id string) (*User, error) {
	// Using DBTX interface methods directly
	row := r.dbTx.QueryRowContext(ctx, `SELECT id, email, password_hash, created_at FROM users WHERE id=$1`, id)
	u := &User{}
	err := row.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return u, nil
}
