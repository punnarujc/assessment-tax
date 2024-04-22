package main

import (
	"github.com/punnarujc/assessment-tax/lib/server"
	"github.com/punnarujc/assessment-tax/router"
	"github.com/shopspring/decimal"
)

func main() {
	decimal.MarshalJSONWithoutQuotes = true

	s := server.New()

	r := router.NewRouter()

	s.WithRouter(r.Router)
	s.Start()
}
