package restaurant

import (
	"context"

	"github.com/Justdan111/swiftEats-backend/internal/db"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}



func (s *Service) ListRestaurants(ctx context.Context) ([]Restaurant, error) {
	rows, err := s.repo.ListRestaurants(ctx)
	if err != nil {
		return nil, err
	}

	var result []Restaurant
	for _, r := range rows {
		result = append(result, s.dbRestaurantToModel(r))
	}

	return result, nil
}

func (s *Service) GetRestaurant(ctx context.Context, id string) (Restaurant, error) {
	r, err := s.repo.GetRestaurantByID(ctx, id)
	if err != nil {
		return Restaurant{}, err
	}

	return s.dbRestaurantToModel(r), nil
}

func (s *Service) GetMenuByRestaurantID(ctx context.Context, restaurantID string) ([]MenuItem, error) {
	rows, err := s.repo.GetMenuItemsByRestaurantID(ctx, restaurantID)
	if err != nil {
		return nil, err
	}

	var items []MenuItem
	for _, m := range rows {
		items = append(items, s.dbMenuItemToModel(m))
	}

	return items, nil
}

//  ADMIN QUERIES 

func (s *Service) CreateRestaurant(ctx context.Context, input RestaurantInput) (Restaurant, error) {
	r, err := s.repo.CreateRestaurant(ctx, input.Name, input.Description, input.Address)
	if err != nil {
		return Restaurant{}, err
	}
	return s.dbRestaurantToModel(r), nil
}

func (s *Service) UpdateRestaurant(ctx context.Context, id string, input RestaurantInput) (Restaurant, error) {
	r, err := s.repo.UpdateRestaurant(ctx, id, input.Name, input.Description, input.Address)
	if err != nil {
		return Restaurant{}, err
	}
	return s.dbRestaurantToModel(r), nil
}

func (s *Service) DeleteRestaurant(ctx context.Context, id string) error {
	return s.repo.DeleteRestaurant(ctx, id)
}

//  MENU ITEM ADMIN QUERIES 

func (s *Service) CreateMenuItem(ctx context.Context, input MenuItemInput) (MenuItem, error) {
	restaurantID := ""
	if input.RestaurantID != nil {
		restaurantID = *input.RestaurantID
	}

	categoryID := ""
	if input.CategoryID != nil {
		categoryID = *input.CategoryID
	}

	m, err := s.repo.CreateMenuItem(ctx, restaurantID, categoryID, input.Name, input.Description, input.PriceCents)
	if err != nil {
		return MenuItem{}, err
	}
	return s.dbMenuItemToModel(m), nil
}

func (s *Service) UpdateMenuItem(ctx context.Context, id string, input MenuItemInput) (MenuItem, error) {
	m, err := s.repo.UpdateMenuItem(ctx, id, input.Name, input.Description, input.PriceCents, input.IsAvailable)
	if err != nil {
		return MenuItem{}, err
	}
	return s.dbMenuItemToModel(m), nil
}

func (s *Service) DeleteMenuItem(ctx context.Context, id string) error {
	return s.repo.DeleteMenuItem(ctx, id)
}

func (s *Service) UpdateMenuItemAvailability(ctx context.Context, id string, isAvailable bool) error {
	return s.repo.UpdateMenuItemAvailability(ctx, id, isAvailable)
}

//  Helpers 

func (s *Service) dbRestaurantToModel(r db.Restaurant) Restaurant {
	return Restaurant{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		Address:     r.Address,
		CreatedAt:   r.CreatedAt,
	}
}

func (s *Service) dbMenuItemToModel(m db.MenuItem) MenuItem {
	return MenuItem{
		ID:           m.ID,
		RestaurantID: m.RestaurantID,
		CategoryID:   m.CategoryID,
		Name:         m.Name,
		Description:  m.Description,
		PriceCents:   m.PriceCents,
		IsAvailable:  m.IsAvailable,
		CreatedAt:    m.CreatedAt,
	}
}
