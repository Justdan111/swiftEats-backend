package restaurant

import (
	"context"
	"database/sql"

	"github.com/Justdan111/swiftEats-backend/internal/db"
	"github.com/google/uuid"
)

type Repository struct {
	q *db.Queries
}

func NewRepository(q *db.Queries) *Repository {
	return &Repository{q: q}
}

// ============ USER QUERIES (Read Only) ============

// ListRestaurants returns all restaurants
func (r *Repository) ListRestaurants(ctx context.Context) ([]db.Restaurant, error) {
	return r.q.GetAllRestaurants(ctx)
}

// GetRestaurantByID returns a single restaurant
func (r *Repository) GetRestaurantByID(ctx context.Context, id string) (db.Restaurant, error) {
	restaurantID, err := uuid.Parse(id)
	if err != nil {
		return db.Restaurant{}, err
	}
	return r.q.GetRestaurantByID(ctx, restaurantID)
}

// GetMenuItemsByRestaurantID returns menu items for a restaurant
func (r *Repository) GetMenuItemsByRestaurantID(ctx context.Context, restaurantID string) ([]db.MenuItem, error) {
	id, err := uuid.Parse(restaurantID)
	if err != nil {
		return nil, err
	}
	return r.q.GetMenuItemsByRestaurantID(ctx, uuid.NullUUID{UUID: id, Valid: true})
}

// ============ ADMIN QUERIES (Create, Update, Delete) ============

// CreateRestaurant creates a new restaurant (admin only)
func (r *Repository) CreateRestaurant(ctx context.Context, name, description, address string) (db.Restaurant, error) {
	return r.q.CreateRestaurant(ctx, db.CreateRestaurantParams{
		Name:        name,
		Description: sql.NullString{String: description, Valid: description != ""},
		Address:     sql.NullString{String: address, Valid: address != ""},
	})
}

// UpdateRestaurant updates an existing restaurant (admin only)
func (r *Repository) UpdateRestaurant(ctx context.Context, id, name, description, address string) (db.Restaurant, error) {
	restaurantID, err := uuid.Parse(id)
	if err != nil {
		return db.Restaurant{}, err
	}

	return r.q.UpdateRestaurant(ctx, db.UpdateRestaurantParams{
		ID:          restaurantID,
		Name:        name,
		Description: sql.NullString{String: description, Valid: description != ""},
		Address:     sql.NullString{String: address, Valid: address != ""},
	})
}

// DeleteRestaurant deletes a restaurant (admin only)
func (r *Repository) DeleteRestaurant(ctx context.Context, id string) error {
	restaurantID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return r.q.DeleteRestaurant(ctx, restaurantID)
}

// ============ MENU ITEM ADMIN QUERIES ============

// CreateMenuItem creates a new menu item (admin only)
func (r *Repository) CreateMenuItem(ctx context.Context, restaurantID, categoryID, name, description string, priceCents int32) (db.MenuItem, error) {
	resID, err := uuid.Parse(restaurantID)
	if err != nil {
		return db.MenuItem{}, err
	}

	catID := uuid.NullUUID{}
	if categoryID != "" {
		parsedCatID, err := uuid.Parse(categoryID)
		if err == nil {
			catID = uuid.NullUUID{UUID: parsedCatID, Valid: true}
		}
	}

	return r.q.CreateMenuItem(ctx, db.CreateMenuItemParams{
		RestaurantID: uuid.NullUUID{UUID: resID, Valid: true},
		CategoryID:   catID,
		Name:         name,
		Description:  sql.NullString{String: description, Valid: description != ""},
		PriceCents:   priceCents,
	})
}

// UpdateMenuItem updates a menu item (admin only)
func (r *Repository) UpdateMenuItem(ctx context.Context, id, name, description string, priceCents int32, isAvailable bool) (db.MenuItem, error) {
	menuItemID, err := uuid.Parse(id)
	if err != nil {
		return db.MenuItem{}, err
	}

	return r.q.UpdateMenuItem(ctx, db.UpdateMenuItemParams{
		ID:          menuItemID,
		Name:        name,
		Description: sql.NullString{String: description, Valid: description != ""},
		PriceCents:  priceCents,
		IsAvailable: sql.NullBool{Bool: isAvailable, Valid: true},
	})
}

// DeleteMenuItem deletes a menu item (admin only)
func (r *Repository) DeleteMenuItem(ctx context.Context, id string) error {
	menuItemID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return r.q.DeleteMenuItem(ctx, menuItemID)
}

// UpdateMenuItemAvailability toggles item availability (admin)
func (r *Repository) UpdateMenuItemAvailability(ctx context.Context, id string, isAvailable bool) error {
	menuItemID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return r.q.UpdateMenuItemAvailability(ctx, db.UpdateMenuItemAvailabilityParams{
		ID:          menuItemID,
		IsAvailable: sql.NullBool{Bool: isAvailable, Valid: true},
	})
}
