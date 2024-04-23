package calculations

import (
	"context"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func Test_Service(t *testing.T) {
	svc := NewService()
	req := Request{
		TotalIncome: decimal.NewFromFloat(500000),
		Wht:         decimal.Zero,
		Allowances:  []Allowance{},
	}

	tax, err := svc.Process(context.Background(), req)

	assert.NoError(t, err)
	assert.True(t, tax.Tax.Equal(decimal.NewFromFloat(29000)), "tax %v should be equal to 29000", tax)
}

func Test_ServiceWithHoldingTax(t *testing.T) {
	svc := NewService()
	req := Request{
		TotalIncome: decimal.NewFromFloat(500000),
		Wht:         decimal.NewFromFloat(25000),
		Allowances:  []Allowance{},
	}

	tax, err := svc.Process(context.Background(), req)

	assert.NoError(t, err)
	assert.True(t, tax.Tax.Equal(decimal.NewFromFloat(4000)), "tax %v should be equal to 4000", tax)
}

func Test_CalculateProgressiveTax(t *testing.T) {
	svc := &serviceImpl{}

	progressiveTax := svc.calculateProgressiveTax(decimal.NewFromFloat(440000))

	assert.True(t, progressiveTax.Equal(decimal.NewFromFloat(29000)), "progressiveTax %v should be equal to 29000", progressiveTax)
}

func Test_CalculateProgressiveTax2M(t *testing.T) {
	svc := &serviceImpl{}

	progressiveTax := svc.calculateProgressiveTax(decimal.NewFromFloat(10000000))

	assert.True(t, progressiveTax.Equal(decimal.NewFromFloat(3110000)), "progressiveTax %v should be equal to 3110000", progressiveTax)
}

func Test_CalculateTotalTaxableAmountOverLimit(t *testing.T) {
	svc := &serviceImpl{}
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

	taxableAmount := svc.calculateTotalTaxableAmount(decimal.NewFromFloat(500000), req)

	assert.True(t, taxableAmount.Equal(decimal.NewFromFloat(340000)), "taxableAmount %v should be equal to 340000", taxableAmount)
}

func Test_CalculateTotalTaxableAmount(t *testing.T) {
	svc := &serviceImpl{}
	req := Request{
		TotalIncome: decimal.NewFromFloat(500000),
		Wht:         decimal.Zero,
		Allowances: []Allowance{
			{
				AllowanceType: "donation",
				Amount:        decimal.NewFromFloat(50000),
			},
		},
	}

	taxableAmount := svc.calculateTotalTaxableAmount(decimal.NewFromFloat(500000), req)

	assert.True(t, taxableAmount.Equal(decimal.NewFromFloat(390000)), "taxableAmount %v should be equal to 390000", taxableAmount)
}
