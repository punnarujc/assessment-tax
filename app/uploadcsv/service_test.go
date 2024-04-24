package uploadcsv

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

// func Test_Service(t *testing.T) {
// 	mockRepo := NewMockRepository(t)
// 	mockRepo.On("GetMaximumDeduction").Return([]TblMaximumDeduction{}, nil)
// 	svc := NewService(mockRepo)
// 	req := Request{
// 		TotalIncome: decimal.NewFromFloat(500000),
// 		Wht:         decimal.Zero,
// 		Allowances:  []Allowance{},
// 	}

// 	tax, err := svc.Process(context.Background(), req)

// 	assert.NoError(t, err)
// 	assert.True(t, tax.Tax.Equal(decimal.NewFromFloat(29000)), "tax %v should be equal to 29000", tax.Tax)
// }

// func Test_ServiceWithHoldingTax(t *testing.T) {
// 	mockRepo := NewMockRepository(t)
// 	mockRepo.On("GetMaximumDeduction").Return([]TblMaximumDeduction{}, nil)
// 	svc := NewService(mockRepo)
// 	req := Request{
// 		TotalIncome: decimal.NewFromFloat(500000),
// 		Wht:         decimal.NewFromFloat(25000),
// 		Allowances:  []Allowance{},
// 	}

// 	tax, err := svc.Process(context.Background(), req)

// 	assert.NoError(t, err)
// 	assert.True(t, tax.Tax.Equal(decimal.NewFromFloat(4000)), "tax %v should be equal to 4000", tax.Tax)
// }

func Test_CalculateProgressiveTax(t *testing.T) {
	svc := &serviceImpl{}

	pt := svc.calculateProgressiveTax(decimal.NewFromFloat(440000))
	pt2m := svc.calculateProgressiveTax(decimal.NewFromFloat(10000000))

	assert.True(t, pt.Equal(decimal.NewFromFloat(29000)), "progressiveTax %v should be equal to 29000", pt)
	assert.True(t, pt2m.Equal(decimal.NewFromFloat(3110000)), "progressiveTax %v should be equal to 3110000", pt2m)
}

func Test_CalculateTotalTaxableAmountOverLimit(t *testing.T) {
	svc := &serviceImpl{}

	taxableAmount := svc.calculateTotalTaxableAmount(decimal.NewFromFloat(500000), decimal.NewFromFloat(200000), []TblMaximumDeduction{})

	assert.True(t, taxableAmount.Equal(decimal.NewFromFloat(340000)), "taxableAmount %v should be equal to 340000", taxableAmount)
}

func Test_CalculateTotalTaxableAmount(t *testing.T) {
	svc := &serviceImpl{}
	tblMaximumDeduction := []TblMaximumDeduction{
		{
			AllowanceType: "personal",
			Amount:        decimal.NewFromFloat(40000),
		},
	}

	taxableAmount := svc.calculateTotalTaxableAmount(decimal.NewFromFloat(500000), decimal.NewFromFloat(50000), []TblMaximumDeduction{})
	taxableAmountWithDb := svc.calculateTotalTaxableAmount(decimal.NewFromFloat(500000), decimal.NewFromFloat(50000), tblMaximumDeduction)

	assert.True(t, taxableAmount.Equal(decimal.NewFromFloat(390000)), "taxableAmount %v should be equal to 390000", taxableAmount)
	assert.True(t, taxableAmountWithDb.Equal(decimal.NewFromFloat(410000)), "taxableAmountWithDb %v should be equal to 3410000", taxableAmountWithDb)
}
