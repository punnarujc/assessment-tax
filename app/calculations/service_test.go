package calculations

import (
	"context"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func Test_Service(t *testing.T) {
	mockRepo := NewMockRepository(t)
	mockRepo.On("GetMaximumDeduction").Return([]TblMaximumDeduction{}, nil)
	svc := NewService(mockRepo)
	req := Request{
		TotalIncome: decimal.NewFromFloat(500000),
		Wht:         decimal.Zero,
		Allowances:  []Allowance{},
	}

	tax, err := svc.Process(context.Background(), req)

	assert.NoError(t, err)
	assert.True(t, tax.Tax.Equal(decimal.NewFromFloat(29000)), "tax %v should be equal to 29000", tax.Tax)
}

func Test_ServiceWithHoldingTax(t *testing.T) {
	mockRepo := NewMockRepository(t)
	mockRepo.On("GetMaximumDeduction").Return([]TblMaximumDeduction{}, nil)
	svc := NewService(mockRepo)
	req := Request{
		TotalIncome: decimal.NewFromFloat(500000),
		Wht:         decimal.NewFromFloat(25000),
		Allowances:  []Allowance{},
	}

	tax, err := svc.Process(context.Background(), req)

	assert.NoError(t, err)
	assert.True(t, tax.Tax.Equal(decimal.NewFromFloat(4000)), "tax %v should be equal to 4000", tax.Tax)
}

func Test_CalculateProgressiveTax(t *testing.T) {
	svc := &serviceImpl{}

	pt, _ := svc.calculateProgressiveTax(decimal.NewFromFloat(440000))
	pt2m, _ := svc.calculateProgressiveTax(decimal.NewFromFloat(10000000))

	assert.True(t, pt.Equal(decimal.NewFromFloat(29000)), "progressiveTax %v should be equal to 29000", pt)
	assert.True(t, pt2m.Equal(decimal.NewFromFloat(3110000)), "progressiveTax %v should be equal to 3110000", pt2m)
}

func Test_CalculateTaxLevel(t *testing.T) {
	svc := &serviceImpl{}

	_, tl := svc.calculateProgressiveTax(decimal.NewFromFloat(340000))
	_, tl2 := svc.calculateProgressiveTax(decimal.NewFromFloat(10000000))

	assert.True(t, tl[0].Tax.Equal(decimal.Zero), "taxLevel %v should be equal to 0", tl[0].Tax)
	assert.True(t, tl[1].Tax.Equal(decimal.NewFromFloat(19000)), "taxLevel %v should be equal to 19000", tl[1].Tax)
	assert.True(t, tl[2].Tax.Equal(decimal.Zero), "taxLevel %v should be equal to 0", tl[2].Tax)
	assert.True(t, tl[3].Tax.Equal(decimal.Zero), "taxLevel %v should be equal to 0", tl[3].Tax)
	assert.True(t, tl[4].Tax.Equal(decimal.Zero), "taxLevel %v should be equal to 0", tl[4].Tax)

	assert.True(t, tl2[0].Tax.Equal(decimal.Zero), "taxLevel %v should be equal to 0", tl2[0].Tax)
	assert.True(t, tl2[1].Tax.Equal(decimal.NewFromFloat(35000)), "taxLevel %v should be equal to 35000", tl2[1].Tax)
	assert.True(t, tl2[2].Tax.Equal(decimal.NewFromFloat(75000)), "taxLevel %v should be equal to 75000", tl2[2].Tax)
	assert.True(t, tl2[3].Tax.Equal(decimal.NewFromFloat(200000)), "taxLevel %v should be equal to 200000", tl2[3].Tax)
	assert.True(t, tl2[4].Tax.Equal(decimal.NewFromFloat(2800000)), "taxLevel %v should be equal to 2800000", tl2[4].Tax)
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

	taxableAmount := svc.calculateTotalTaxableAmount(decimal.NewFromFloat(500000), req, []TblMaximumDeduction{})

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
			{
				AllowanceType: "k-receipt",
				Amount:        decimal.NewFromFloat(200000),
			},
		},
	}
	tblMaximumDeduction := []TblMaximumDeduction{
		{
			AllowanceType: "personal",
			Amount:        decimal.NewFromFloat(40000),
		},
		{
			AllowanceType: "k-receipt",
			Amount:        decimal.NewFromFloat(60000),
		},
	}

	taxableAmount := svc.calculateTotalTaxableAmount(decimal.NewFromFloat(500000), req, []TblMaximumDeduction{})
	taxableAmountWithDb := svc.calculateTotalTaxableAmount(decimal.NewFromFloat(500000), req, tblMaximumDeduction)

	assert.True(t, taxableAmount.Equal(decimal.NewFromFloat(340000)), "taxableAmount %v should be equal to 390000", taxableAmount)
	assert.True(t, taxableAmountWithDb.Equal(decimal.NewFromFloat(360000)), "taxableAmountWithDb %v should be equal to 3600000", taxableAmountWithDb)
}
