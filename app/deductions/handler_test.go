package deductions

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func Test_AllowanceType(t *testing.T) {
	var req = &Request{}

	personal := req.isAllowanceTypeValid("personal")
	kReceipt := req.isAllowanceTypeValid("k-receipt")
	insurance := req.isAllowanceTypeValid("insurance")

	assert.True(t, personal, "personal should be valid")
	assert.True(t, kReceipt, "k-receipt should be valid")
	assert.False(t, insurance, "insurance should not be valid")
}

func Test_AmountOverLimit(t *testing.T) {
	var req = &Request{
		Amount: decimal.NewFromFloat(110000),
	}

	receipt := req.isAmountValid("k-receipt")
	personal := req.isAmountValid("personal")

	assert.False(t, receipt, "receipt should not be valid")
	assert.False(t, personal, "personal should not be valid")
}

func Test_AmountUnderLimit(t *testing.T) {
	var req = &Request{
		Amount: decimal.NewFromFloat(5000),
	}

	receipt := req.isAmountValid("k-receipt")
	personal := req.isAmountValid("personal")

	assert.True(t, receipt, "receipt should be valid")
	assert.False(t, personal, "personal should not be valid")
}
