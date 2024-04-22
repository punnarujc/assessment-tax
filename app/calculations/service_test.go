package calculations

import (
	"context"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func Test_CalculateTotalTaxableAmount(t *testing.T) {
	svc := NewService()

	req := Request{
		TotalIncome: decimal.NewFromFloat(500000),
		Wht:         decimal.Zero,
		Allowances:  []Allowance{},
	}

	tax, err := svc.Process(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, tax.Tax.String(), decimal.NewFromFloat(29000).String())
}

func Test_CalculateTotalTaxableAmountMoreThan2M(t *testing.T) {
	svc := NewService()

	req := Request{
		TotalIncome: decimal.NewFromFloat(10000000),
		Wht:         decimal.Zero,
		Allowances:  []Allowance{},
	}

	tax, err := svc.Process(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, tax.Tax.String(), decimal.NewFromFloat(3089000).String())
}

func Test_CalculateTotalTaxableAmountWithAllowances(t *testing.T) {
	svc := NewService()

	req := Request{
		TotalIncome: decimal.NewFromFloat(500000),
		Wht:         decimal.Zero,
		Allowances: []Allowance{
			{
				AllowanceType: "donation",
				Amount:        decimal.NewFromFloat(200000),
			},
		},
	}

	tax, err := svc.Process(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, tax.Tax.String(), decimal.NewFromFloat(19000).String())
}
