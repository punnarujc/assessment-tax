package uploadcsv

import "github.com/shopspring/decimal"

type Response struct {
	Taxes []Tax `json:"taxes"`
}

type Tax struct {
	TotalIncome decimal.Decimal `json:"totalIncome"`
	Tax         decimal.Decimal `json:"tax"`
	TaxRefund   decimal.Decimal `json:"taxRefund"`
}

type ProgressiveTaxRatio struct {
	TaxLevel   string
	Amount     decimal.Decimal
	TaxPercent decimal.Decimal
}

type TaxFile struct {
	TotalIncome decimal.Decimal
	Wht         decimal.Decimal
	Donation    decimal.Decimal
}

type Allowance struct {
	AllowanceType string          `json:"allowanceType"`
	Amount        decimal.Decimal `json:"amount"`
}
