package uploadcsv

import "gorm.io/gorm"

//go:generate mockery --dir=. --name=Repository  --output=. --filename=repository_mock.go --structname=MockRepository --outpkg=uploadcsv --inpackage=true
type Repository interface {
	GetMaximumDeduction() ([]TblMaximumDeduction, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db: db,
	}
}

func (r *repositoryImpl) GetMaximumDeduction() ([]TblMaximumDeduction, error) {
	var tblMaximumDeductionList []TblMaximumDeduction

	err := r.db.Debug().Model(&TblMaximumDeduction{}).
		Find(&tblMaximumDeductionList).Error

	return tblMaximumDeductionList, err
}
