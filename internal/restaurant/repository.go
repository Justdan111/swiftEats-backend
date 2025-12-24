package restaurant

import (
	"context"

	"github.com/Justdan111/swiftEats-backend/internal/db"
)

type Repository struct {
	q *db.Queries
}

func NewRepository(q *db.Queries) *Repository {
	return &Repository{q: q}
}

// get all restaurants
func (r *Repository) ListRestaurants(ctx context.Context) ([]db.Restaurant, error) {
	return r.q.ListRestaurants(ctx)
}

// get one restaurant by 
func (r *Repository) GetRestaurant(ctx context.Context, id string) (db.Restaurant, error) {
	return r.q.GetRestaurantByID(ctx, id)
}

// Get menu for restaurant
func (r *Repository) GetMenuItemsByRestaurantID(ctx context.Context, restaurantID string) ([]db.MenuItem, error) {
	return r.q.GetMenuItemsByRestaurantID(ctx, restaurantID)
}

// Add a new restaurant
func (r *Repository) CreateRestaurant(ctx context.Context, name, address string, isOpen bool) (db.Restaurant, error) {
	return r.q.CreateRestaurant(ctx, db.CreateRestaurantParams{
		Name:    name,
		Address: address,
		IsOpen:  isOpen,
	})
}

// Update restaurant details
func (r *Repository) UpdateRestaurant(ctx context.Context, id, name, address string, isOpen bool) (db.Restaurant, error) {
	return r.q.UpdateRestaurant(ctx, db.UpdateRestaurantParams{
		ID:      id,
		Name:    name,
		Address: address,
		IsOpen:  isOpen,
	})
}

// Delete a restaurant
func (r *Repository) DeleteRestaurant(ctx context.Context, id string) error {
	return r.q.DeleteRestaurant(ctx, id)
}

// Add a new menu item
func (r *Repository) CreateMenuItem(ctx context.Context, restaurantID, name, description string, price float64) (db.MenuItem, error) {
	return r.q.CreateMenuItem(ctx, db.CreateMenuItemParams{
		RestaurantID: restaurantID,
		Name:         name,
		Description:  description,
		Price:        price,
	})
}

// Update a menu item
func (r *Repository) UpdateMenuItem(ctx context.Context, id, name, description string, price float64) (db.MenuItem, error) {
	return r.q.UpdateMenuItem(ctx, db.UpdateMenuItemParams{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
	})
}

// Delete a menu item
func (r *Repository) DeleteMenuItem(ctx context.Context, id string) error {
	return r.q.DeleteMenuItem(ctx, id)
}	



