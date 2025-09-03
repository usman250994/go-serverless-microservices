package user

import "context"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) updateProfile(ctx context.Context, r profile) (string, error) {

	if err := s.repo.Save(ctx, &user); err != nil {
		return "", err
	}

	return "success", nil
}
