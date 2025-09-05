package item

import (
	"context"
)

type Service struct {
	repo *Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: &repo}
}

func (s *Service) updateProfile(ctx context.Context, p profile) (string, error) {
	res, err := s.repo.Save(ctx, &p)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (s *Service) getProfile(ctx context.Context, userId string) (*profile, error) {
	profile, err := s.repo.FindById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return profile, nil
}
