package cart

import "time"

type Cart struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	RestaurantID string    `json:"restaurant_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CartItem struct {
	ID 	  string    `json:"id"`
	CartID   string    `json:"cart_id"`
	ManueItemID string    `json:"menu_item_id"`
	Quantity int 	 `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}