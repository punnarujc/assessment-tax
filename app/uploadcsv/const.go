package uploadcsv

import "github.com/shopspring/decimal"

const (
	PARAM_TAX_FILE string = "taxFile"
	CSV_DELIMITER  string = ","

	TAX_LEVEL_150K    string = "0-150,000"
	TAX_LEVEL_500K    string = "150,001-500,000"
	TAX_LEVEL_1M      string = "500,001-1,000,000"
	TAX_LEVEL_2M      string = "1,000,000-2,000,000"
	TAX_LEVEL_MORE_2M string = "2,000,000 ขึ้นไป"

	ALLOWANCE_TYPE_DONATION  string = "donation"
	ALLOWANCE_TYPE_K_RECEIPT string = "k-receipt"
	ALLOWANCE_TYPE_PERSONAL  string = "personal"
)

var (
	PERSONAL_DEDUCTION_60K decimal.Decimal = decimal.NewFromFloat(60000)

	TAX_AMOUNT_150K decimal.Decimal = decimal.NewFromInt(150000)
	TAX_AMOUNT_350K decimal.Decimal = decimal.NewFromInt(350000)
	TAX_AMOUNT_500K decimal.Decimal = decimal.NewFromInt(500000)
	TAX_AMOUNT_1M   decimal.Decimal = decimal.NewFromInt(1000000)

	PROGRESSIVE_TAX_RATIO []ProgressiveTaxRatio = []ProgressiveTaxRatio{
		{
			TaxLevel:   TAX_LEVEL_150K,
			Amount:     TAX_AMOUNT_150K,
			TaxPercent: decimal.Zero,
		},
		{
			TaxLevel:   TAX_LEVEL_500K,
			Amount:     TAX_AMOUNT_350K,
			TaxPercent: decimal.NewFromFloat(0.1),
		},
		{
			TaxLevel:   TAX_LEVEL_1M,
			Amount:     TAX_AMOUNT_500K,
			TaxPercent: decimal.NewFromFloat(0.15),
		},
		{
			TaxLevel:   TAX_LEVEL_2M,
			Amount:     TAX_AMOUNT_1M,
			TaxPercent: decimal.NewFromFloat(0.2),
		},
		{
			TaxLevel:   TAX_LEVEL_MORE_2M,
			TaxPercent: decimal.NewFromFloat(0.35),
		},
	}

	ALLOWANCE_MAX_AMOUNT map[string]decimal.Decimal = map[string]decimal.Decimal{
		ALLOWANCE_TYPE_DONATION:  decimal.NewFromInt(100000),
		ALLOWANCE_TYPE_K_RECEIPT: decimal.NewFromInt(50000),
		ALLOWANCE_TYPE_PERSONAL:  decimal.NewFromInt(100000),
	}
)
