package item

import (
	"context"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) addProduct(ctx context.Context, p product) (string, error) {
	_, err := s.repo.Save(ctx, &p)
	if err != nil {
		return "", err
	}

	return "success", nil
}

func (s *Service) getNearestProducts(ctx context.Context, q *ProductQuery) ([]product, error) {
	prd, err := s.repo.getNearestTen(ctx, q)
	if err != nil {
		return nil, err
	}
	return prd, nil
}
