package main

import (
	"github.com/punnarujc/assessment-tax/lib/server"
	"github.com/punnarujc/assessment-tax/router"
)

func main() {
	s := server.New()

	r := router.NewRouter()

	s.WithRouter(r.Router)
	s.Start()
}
