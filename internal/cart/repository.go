package cart

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetCartByUser(userID string) (*Cart, error) {
	var c Cart
	err := r.db.QueryRow(`
		SELECT id, user_id, restaurant_id, created_at, updated_at
		FROM carts WHERE user_id = $1
	`, userID).Scan(&c.ID, &c.UserID, &c.RestaurantID, &c.CreatedAt, &c.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &c, err
}

func (r *Repository) CreateCart(userID, restaurantID string) (*Cart, error) {
	var c Cart
	err := r.db.QueryRow(`
		INSERT INTO carts (user_id, restaurant_id)
		VALUES ($1, $2)
		RETURNING id, user_id, restaurant_id, created_at, updated_at
	`, userID, restaurantID).Scan(&c.ID, &c.UserID, &c.RestaurantID, &c.CreatedAt, &c.UpdatedAt)

	return &c, err
}

func (r *Repository) AddItem(cartID, menuItemID string, qty int) error {
	_, err := r.db.Exec(`
		INSERT INTO cart_items (cart_id, menu_item_id, quantity)
		VALUES ($1, $2, $3)
		ON CONFLICT (cart_id, menu_item_id)
		DO UPDATE SET quantity = cart_items.quantity + $3, updated_at = now()
	`, cartID, menuItemID, qty)

	return err
}
