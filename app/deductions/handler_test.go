package deductions

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func Test_AllowanceType(t *testing.T) {
	var req = &Request{}

	personal := req.isAllowanceTypeValid("personal")
	insurance := req.isAllowanceTypeValid("insurance")

	assert.True(t, personal, "personal should be valid")
	assert.False(t, insurance, "insurance should not be valid")
}

func Test_AmountOverLimit(t *testing.T) {
	var req = &Request{
		Amount: decimal.NewFromFloat(70000),
	}

	donation := req.isAmountValid("donation")
	personal := req.isAmountValid("personal")

	assert.True(t, donation, "donation should be valid")
	assert.False(t, personal, "personal should not be valid")
}
