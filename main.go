package main

import (
	"github.com/punnarujc/assessment-tax/app/deductions"
	"github.com/punnarujc/assessment-tax/lib/postgresql"
	"github.com/punnarujc/assessment-tax/lib/server"
	"github.com/punnarujc/assessment-tax/router"
	"github.com/shopspring/decimal"
)

func main() {
	decimal.MarshalJSONWithoutQuotes = true

	s := server.New()
	db := postgresql.New()

	err := db.AutoMigrate(&deductions.TblMaximumDeduction{})
	if err != nil {
		panic("AutoMigrate failed")
	}

	r := router.NewRouter(db)

	s.WithRouter(r.Router)
	s.Start()
}
