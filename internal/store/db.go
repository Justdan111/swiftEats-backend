package store

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	DB *pgxpool.Pool
}

func NewStore(connStr string) (*Store, error) {
	pool , err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to create DB pool: %w", err)
	}

	return &Store{DB: pool}, nil
}

func (s *Store) Close() {
	s.DB.Close()
}