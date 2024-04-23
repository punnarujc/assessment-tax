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
		ALLOWANCE_TYPE_K_RECEIPT: decimal.NewFromFloat(100000),
		ALLOWANCE_TYPE_PERSONAL:  decimal.NewFromFloat(100000),
	}

	ALLOWANCE_MIN_AMOUNT map[string]decimal.Decimal = map[string]decimal.Decimal{
		ALLOWANCE_TYPE_K_RECEIPT: decimal.Zero,
		ALLOWANCE_TYPE_PERSONAL:  decimal.NewFromFloat(10000),
	}
)
