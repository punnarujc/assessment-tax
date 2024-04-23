package calculations

import "github.com/shopspring/decimal"

type TblMaximumDeduction struct {
	AllowanceType string          `gorm:"column:allowance_type"`
	Amount        decimal.Decimal `gorm:"column:amount"`
}
