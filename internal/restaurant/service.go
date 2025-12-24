package restaurant

import "context"

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
		result = append(result, Restaurant{
			ID:        r.ID,
			Name:      r.Name,
			Address:   r.Address,
			IsOpen:    r.IsOpen,
			CreatedAt: r.CreatedAt,
		})
	}

	return result, nil
}

func (s *Service) GetRestaurant(ctx context.Context, id string) (Restaurant, error) {
	r, err := s.repo.GetRestaurant(ctx, id)
	if err != nil {
		return Restaurant{}, err
	}

	return Restaurant{
		ID:        r.ID,
		Name:      r.Name,
		Address:   r.Address,
		IsOpen:    r.IsOpen,
		CreatedAt: r.CreatedAt,
	}, nil
}

func (s *Service) GetMenu(ctx context.Context, restaurantID string) ([]MenuItem, error) {
	rows, err := s.repo.GetMenu(ctx, restaurantID)
	if err != nil {
		return nil, err
	}

	var items []MenuItem
	for _, m := range rows {
		items = append(items, MenuItem{
			ID:           m.ID,
			RestaurantID: m.RestaurantID,
			Name:         m.Name,
			Description:  m.Description,
			Price:        m.Price,
		})
	}

	return items, nil
}
