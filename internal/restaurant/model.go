package restaurant

import (
	"database/sql"

	"github.com/google/uuid"
)


type Restaurant struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Address     sql.NullString `json:"address"`
	CreatedAt   sql.NullTime   `json:"created_at"`
}


type MenuItem struct {
	ID           uuid.UUID      `json:"id"`
	RestaurantID uuid.NullUUID  `json:"restaurant_id"`
	CategoryID   uuid.NullUUID  `json:"category_id"`
	Name         string         `json:"name"`
	Description  sql.NullString `json:"description"`
	PriceCents   int32          `json:"price_cents"` 
	IsAvailable  sql.NullBool   `json:"is_available"`
	CreatedAt    sql.NullTime   `json:"created_at"`
}


type MenuItemInput struct {
	RestaurantID *string `json:"restaurant_id"`
	CategoryID   *string `json:"category_id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	PriceCents   int32   `json:"price_cents"`
	IsAvailable  bool    `json:"is_available"`
}


type RestaurantInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
}
