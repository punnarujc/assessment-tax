package postgresql

import (
	"github.com/punnarujc/assessment-tax/lib/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	dbCfg := config.New().GetDb()

	db, err := gorm.Open(postgres.Open(dbCfg), &gorm.Config{})
	if err != nil {
		panic("Connect to Postgresql failed")
	}

	return db
}
