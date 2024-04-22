package calculations

import (
	"net/http"

	"github.com/punnarujc/assessment-tax/lib/server"
)

type Handler interface {
	Calculations(c server.Context) error
}

type handlerImpl struct {
}

func NewHandler() Handler {
	return &handlerImpl{}
}

func (h *handlerImpl) Calculations(c server.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{})
}
