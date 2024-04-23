package calculations

import (
	"context"

	"github.com/shopspring/decimal"
)

type Service interface {
	Process(ctx context.Context, req Request) (Response, error)
}

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &serviceImpl{
		repo: repo,
	}
}

func (s *serviceImpl) Process(ctx context.Context, req Request) (Response, error) {
	tblMaximumDeduction, err := s.repo.GetMaximumDeduction()
	if err != nil {
		return Response{}, err
	}

	totalTaxable := s.calculateTotalTaxableAmount(req.TotalIncome, req, tblMaximumDeduction)

	progressiveTax, taxLevel := s.calculateProgressiveTax(totalTaxable)

	tax := progressiveTax.Sub(req.Wht)

	var resp = Response{
		Tax:      tax,
		TaxLevel: taxLevel,
	}

	return resp, nil
}

func (s *serviceImpl) calculateTotalTaxableAmount(amount decimal.Decimal, req Request, tblMaximumDeduction []TblMaximumDeduction) decimal.Decimal {
	var taxableAmount = amount

	personalDeduct := PERSONAL_DEDUCTION_60K
	for _, v := range tblMaximumDeduction {
		if v.AllowanceType == ALLOWANCE_TYPE_PERSONAL {
			personalDeduct = v.Amount
		}
	}

	allowances := append(req.Allowances, Allowance{
		AllowanceType: ALLOWANCE_TYPE_PERSONAL,
		Amount:        personalDeduct,
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

func (s *serviceImpl) calculateProgressiveTax(amount decimal.Decimal) (decimal.Decimal, []TaxLevel) {
	var tax decimal.Decimal
	var taxLevel = make([]TaxLevel, 0, len(PROGRESSIVE_TAX_RATIO))

	for _, ptr := range PROGRESSIVE_TAX_RATIO {
		var taxPerLevel = TaxLevel{
			Level: ptr.TaxLevel,
			Tax:   decimal.Zero,
		}

		switch {
		case ptr.TaxLevel == TAX_LEVEL_MORE_2M:
			tax = tax.Add(amount.Mul(ptr.TaxPercent))
			taxPerLevel.Tax = amount.Mul(ptr.TaxPercent)

		case amount.GreaterThan(ptr.Amount):
			tax = tax.Add(ptr.Amount.Mul(ptr.TaxPercent))
			taxPerLevel.Tax = ptr.Amount.Mul(ptr.TaxPercent)
			amount = amount.Sub(ptr.Amount)

		case amount.GreaterThan(decimal.Zero):
			tax = tax.Add(amount.Mul(ptr.TaxPercent))
			taxPerLevel.Tax = amount.Mul(ptr.TaxPercent)
			amount = decimal.Zero

		default:
		}

		taxLevel = append(taxLevel, taxPerLevel)
	}

	return tax, taxLevel
}
