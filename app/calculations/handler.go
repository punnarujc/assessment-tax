package calculations

import (
	"net/http"

	"github.com/punnarujc/assessment-tax/lib/errs"
	"github.com/punnarujc/assessment-tax/lib/server"
)

type Handler interface {
	Calculations(c server.Context) error
}

type handlerImpl struct {
	svc Service
}

func NewHandler(svc Service) Handler {
	return &handlerImpl{
		svc: svc,
	}
}

func (h *handlerImpl) Calculations(c server.Context) error {
	var req Request

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errs.NewErrContextWithError(errs.BadRequest, err))
	}

	resp, err := h.svc.Process(c.Request().Context(), req)
	if err != nil {
		return c.JSON(errs.GenericError.HttpStatus, err)
	}

	return c.JSON(http.StatusOK, resp)
}
