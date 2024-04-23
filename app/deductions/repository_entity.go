package deductions

import "github.com/shopspring/decimal"

type TblMaximumDeduction struct {
	AllowanceType string          `gorm:"column:allowance_type;primaryKey"`
	Amount        decimal.Decimal `gorm:"column:amount;type:decimal(8,2)"`
}
