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

	var resp = Response{}

	switch allowanceType {
	case ALLOWANCE_TYPE_PERSONAL:
		resp.PersonalDeduction = &req.Amount
	case ALLOWANCE_TYPE_K_RECEIPT:
		resp.Kreceipt = &req.Amount
	default:
	}

	return resp, nil
}
