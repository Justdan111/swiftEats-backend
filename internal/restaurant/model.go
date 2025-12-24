package restaurant

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Restaurant represents a restaurant (read for user/admin)
type Restaurant struct {
	ID          uuid.UUID       `json:"id"`
	Name        string          `json:"name"`
	Description sql.NullString  `json:"description"`
	Address     string          `json:"address"`
	CreatedAt   sql.NullTime    `json:"created_at"`
}

// MenuItem represents a menu item with price in cents
type MenuItem struct {
	ID           uuid.UUID      `json:"id"`
	RestaurantID uuid.NullUUID  `json:"restaurant_id"`
	CategoryID   sql.NullString `json:"category_id"`
	Name         string         `json:"name"`
	Description  sql.NullString `json:"description"`
	PriceCents   int32          `json:"price_cents"` // In cents, not float
	IsAvailable  sql.NullBool   `json:"is_available"`
	CreatedAt    sql.NullTime   `json:"created_at"`
}

// MenuItemInput is for admin creating/updating menu items
type MenuItemInput struct {
	RestaurantID *string `json:"restaurant_id"`
	CategoryID   *string `json:"category_id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	PriceCents   int32   `json:"price_cents"`
	IsAvailable  bool    `json:"is_available"`
}

// RestaurantInput is for admin creating/updating restaurants
type RestaurantInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
}