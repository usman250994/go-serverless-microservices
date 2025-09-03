package user

import "context"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) updateProfile(ctx context.Context, p profile) (string, error) {

	res, err := s.repo.Save(ctx, &p)
	if err != nil {
		return "", err
	}

	return res, nil
}
