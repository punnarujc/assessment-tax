package router

import (
	"github.com/punnarujc/assessment-tax/app/calculations"
	"github.com/punnarujc/assessment-tax/lib/server"
)

type Router interface {
	Router(server.EchoServer)
}

type routerImpl struct {
}

func NewRouter() Router {
	return &routerImpl{}
}

func (r *routerImpl) Router(s server.EchoServer) {
	calculationsHandler := calculations.New()

	s.POST("/tax/calculations", calculationsHandler.Calculations)
}
