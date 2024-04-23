package deductions

import (
	"context"
)

type Service interface {
	Process(ctx context.Context, req Request, allowanceType string) (Response, error)
}

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &serviceImpl{
		repo: repo,
	}
}

func (s *serviceImpl) Process(ctx context.Context, req Request, allowanceType string) (Response, error) {
	err := s.repo.UpsertMaximumDeduction(allowanceType, req.Amount)
	if err != nil {
		return Response{}, err
	}

	var resp = Response{
		PersonalDeduction: req.Amount,
	}

	return resp, nil
}
