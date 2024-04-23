package deductions

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

//go:generate mockery --dir=. --name=Repository  --output=. --filename=repository_mock.go --structname=MockRepository --outpkg=deductions --inpackage=true
type Repository interface {
	// GetMaximumDeduction(allowanceType string) (TblMaximumDeduction, error)
	UpsertMaximumDeduction(allowanceType string, amount decimal.Decimal) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db: db,
	}
}

// func (r *repositoryImpl) GetMaximumDeduction(allowanceType string) (TblMaximumDeduction, error) {
// 	var tblMaximumDeduction TblMaximumDeduction

// 	err := r.db.Debug().Model(&TblMaximumDeduction{}).
// 		Where("allowance_type = ?", allowanceType).
// 		Take(&tblMaximumDeduction)

// 	if err != nil {
// 		return TblMaximumDeduction{}, err.Error
// 	}

// 	return tblMaximumDeduction, nil
// }

func (r *repositoryImpl) UpsertMaximumDeduction(allowanceType string, amount decimal.Decimal) error {
	var newDeduction = TblMaximumDeduction{
		AllowanceType: allowanceType,
		Amount:        amount,
	}

	return r.db.Debug().Save(&newDeduction).Error
}
