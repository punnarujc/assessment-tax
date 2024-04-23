package deductions

import "gorm.io/gorm"

func New(db *gorm.DB) Handler {
	repo := NewRepository(db)
	svc := NewService(repo)

	return NewHandler(svc)
}
