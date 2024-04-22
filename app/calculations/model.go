package calculations

import "github.com/shopspring/decimal"

type Request struct {
	TotalIncome decimal.Decimal `json:"totalIncome"`
	Wht         decimal.Decimal `json:"wht"`
	Allowances  []Allowance     `json:"allowances"`
}

type Allowance struct {
	AllowanceType string          `json:"allowanceType"`
	Amount        decimal.Decimal `json:"amount"`
}

type Response struct {
	Tax decimal.Decimal `json:"tax"`
}

type ProgressiveTaxRatio struct {
	TaxLevel   string
	Amount     decimal.Decimal
	TaxPercent decimal.Decimal
}
