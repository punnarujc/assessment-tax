package deductions

import "github.com/shopspring/decimal"

type Request struct {
	Amount decimal.Decimal `json:"amount"`
}

func (r *Request) isAllowanceTypeValid(allowanceType string) bool {
	_, ok := ALLOWANCE_MAX_AMOUNT[allowanceType]
	return ok
}

func (r *Request) isAmountValid(allowanceType string) bool {
	return ALLOWANCE_MAX_AMOUNT[allowanceType].GreaterThan(r.Amount)
}

type Response struct {
	PersonalDeduction decimal.Decimal `json:"personalDeduction"`
}
