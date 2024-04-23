package router

import (
	"github.com/punnarujc/assessment-tax/app/calculations"
	"github.com/punnarujc/assessment-tax/app/deductions"
	"github.com/punnarujc/assessment-tax/lib/server"
	"gorm.io/gorm"
)

type Router interface {
	Router(server.EchoServer)
}

type routerImpl struct {
	db *gorm.DB
}

func NewRouter(db *gorm.DB) Router {
	return &routerImpl{
		db: db,
	}
}

func (r *routerImpl) Router(s server.EchoServer) {
	calculationsHandler := calculations.New(r.db)
	deductionsHandler := deductions.New(r.db)

	s.POST("/tax/calculations", calculationsHandler.Calculations)
	s.POST("/admin/deductions/:allowanceType", deductionsHandler.Deductions)
}
