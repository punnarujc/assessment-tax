package uploadcsv

import (
	"bufio"
	"context"
	"mime/multipart"
	"strings"

	"github.com/shopspring/decimal"
)

type Service interface {
	Process(ctx context.Context, taxFile *multipart.FileHeader) (Response, error)
}

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &serviceImpl{
		repo: repo,
	}
}

func (s *serviceImpl) Process(ctx context.Context, taxFile *multipart.FileHeader) (Response, error) {
	tblMaximumDeduction, err := s.repo.GetMaximumDeduction()
	if err != nil {
		return Response{}, err
	}

	taxFileList, err := s.mappingTaxDetail(taxFile)
	if err != nil {
		return Response{}, err
	}

	taxes := make([]Tax, 0, len(taxFileList))

	for _, tf := range taxFileList {
		totalTaxable := s.calculateTotalTaxableAmount(tf.TotalIncome, tf.Donation, tblMaximumDeduction)

		progressiveTax := s.calculateProgressiveTax(totalTaxable)
		totalTax := progressiveTax.Sub(tf.Wht)
		taxRefund := decimal.Zero
		if totalTax.LessThan(decimal.Zero) {
			totalTax = decimal.Zero
			taxRefund = progressiveTax.Sub(tf.Wht).Neg()
		}

		var tax = Tax{
			TotalIncome: tf.TotalIncome,
			Tax:         totalTax,
			TaxRefund:   taxRefund,
		}

		taxes = append(taxes, tax)
	}

	var resp = Response{
		Taxes: taxes,
	}

	return resp, nil
}

func (s *serviceImpl) mappingTaxDetail(taxFile *multipart.FileHeader) ([]TaxFile, error) {
	var taxFileDetail = make([]TaxFile, 0, 10)

	f, err := taxFile.Open()
	if err != nil {
		return []TaxFile{}, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var i int = 0
	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case i > 0:
			arr := strings.Split(line, CSV_DELIMITER)
			taxFile, err := s.validateTaxFileDetail(arr)
			if err != nil {
				return []TaxFile{}, err
			}

			taxFileDetail = append(taxFileDetail, taxFile)
		default:
		}

		i++
	}

	return taxFileDetail, nil
}

func (s *serviceImpl) validateTaxFileDetail(arr []string) (TaxFile, error) {
	totalIncome, err := decimal.NewFromString(arr[0])
	if err != nil {
		return TaxFile{}, err
	}
	wht, err := decimal.NewFromString(arr[1])
	if err != nil {
		return TaxFile{}, err
	}
	donation, err := decimal.NewFromString(arr[2])
	if err != nil {
		return TaxFile{}, err
	}

	taxFile := TaxFile{
		TotalIncome: totalIncome,
		Wht:         wht,
		Donation:    donation,
	}

	return taxFile, nil
}

func (s *serviceImpl) calculateTotalTaxableAmount(amount decimal.Decimal, donation decimal.Decimal, tblMaximumDeduction []TblMaximumDeduction) decimal.Decimal {
	var allowances = make([]Allowance, 0)
	var taxableAmount = amount

	personalDeduct := PERSONAL_DEDUCTION_60K
	for _, v := range tblMaximumDeduction {
		if v.AllowanceType == ALLOWANCE_TYPE_PERSONAL {
			personalDeduct = v.Amount
		}
	}

	allowances = append(allowances, Allowance{
		AllowanceType: ALLOWANCE_TYPE_PERSONAL,
		Amount:        personalDeduct,
	}, Allowance{
		AllowanceType: ALLOWANCE_TYPE_DONATION,
		Amount:        donation,
	})

	for _, alw := range allowances {
		deductAmount := alw.Amount
		if alw.Amount.GreaterThan(ALLOWANCE_MAX_AMOUNT[alw.AllowanceType]) {
			deductAmount = ALLOWANCE_MAX_AMOUNT[alw.AllowanceType]
		}
		taxableAmount = taxableAmount.Sub(deductAmount)
	}

	return taxableAmount
}

func (s *serviceImpl) calculateProgressiveTax(amount decimal.Decimal) decimal.Decimal {
	var tax decimal.Decimal

	for _, ptr := range PROGRESSIVE_TAX_RATIO {
		switch {
		case ptr.TaxLevel == TAX_LEVEL_MORE_2M:
			tax = tax.Add(amount.Mul(ptr.TaxPercent))

		case amount.GreaterThan(ptr.Amount):
			tax = tax.Add(ptr.Amount.Mul(ptr.TaxPercent))
			amount = amount.Sub(ptr.Amount)

		case amount.GreaterThan(decimal.Zero):
			tax = tax.Add(amount.Mul(ptr.TaxPercent))
			amount = decimal.Zero

		default:
		}
	}

	return tax
}
