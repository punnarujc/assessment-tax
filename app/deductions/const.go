package deductions

import "github.com/shopspring/decimal"

const (
	PARAM_ALLOWANCE_TYPE string = "allowanceType"

	ALLOWANCE_TYPE_DONATION  string = "donation"
	ALLOWANCE_TYPE_K_RECEIPT string = "k-receipt"
	ALLOWANCE_TYPE_PERSONAL  string = "personal"
)

var (
	ALLOWANCE_MAX_AMOUNT map[string]decimal.Decimal = map[string]decimal.Decimal{
		ALLOWANCE_TYPE_DONATION:  decimal.NewFromInt(100000),
		ALLOWANCE_TYPE_K_RECEIPT: decimal.NewFromInt(50000),
		ALLOWANCE_TYPE_PERSONAL:  decimal.NewFromInt(60000),
	}
)
