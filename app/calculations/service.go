package calculations

import (
	"context"

	"github.com/shopspring/decimal"
)

type Service interface {
	Process(ctx context.Context, req Request) (Response, error)
}

type serviceImpl struct {
}

func NewService() Service {
	return &serviceImpl{}
}

func (s *serviceImpl) Process(ctx context.Context, req Request) (Response, error) {
	totalTaxable := s.calculateTotalTaxableAmount(req.TotalIncome, req)

	progressiveTax := s.calculateProgressiveTax(totalTaxable)

	tax := progressiveTax.Sub(req.Wht)

	var resp = Response{
		Tax: tax,
	}

	return resp, nil
}

func (s *serviceImpl) calculateTotalTaxableAmount(amount decimal.Decimal, req Request) decimal.Decimal {
	var taxableAmount = amount

	allowances := append(req.Allowances, Allowance{
		AllowanceType: ALLOWANCE_TYPE_PERSONAL,
		Amount:        PERSONAL_DEDUCTION_60K,
	})

	for _, alw := range allowances {
		deductAmount := alw.Amount
		if alw.Amount.GreaterThan(ALLOWANCE_MAX_AMOUNT[alw.AllowanceType]) {
			deductAmount = ALLOWANCE_MAX_AMOUNT[alw.AllowanceType]
		}
		taxableAmount = taxableAmount.Sub(deductAmount)
	}

	return taxableAmount
}

func (s *serviceImpl) calculateProgressiveTax(amount decimal.Decimal) decimal.Decimal {
	var tax decimal.Decimal

	for _, ptr := range PROGRESSIVE_TAX_RATIO {
		if ptr.TaxLevel == TAX_LEVEL_MORE_2M {
			tax = tax.Add(amount.Mul(ptr.TaxPercent))
			break
		}

		if amount.GreaterThan(ptr.Amount) {
			tax = tax.Add(ptr.Amount.Mul(ptr.TaxPercent))
			amount = amount.Sub(ptr.Amount)
		} else {
			tax = tax.Add(amount.Mul(ptr.TaxPercent))
			break
		}
	}

	return tax
}
