package deductions

import (
	"errors"
	"net/http"

	"github.com/punnarujc/assessment-tax/lib/errs"
	"github.com/punnarujc/assessment-tax/lib/server"
)

type Handler interface {
	Deductions(c server.Context) error
}

type handlerImpl struct {
	svc Service
}

func NewHandler(svc Service) Handler {
	return &handlerImpl{
		svc: svc,
	}
}

func (h *handlerImpl) Deductions(c server.Context) error {
	var req Request
	var allowanceType = c.Param(PARAM_ALLOWANCE_TYPE)

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errs.NewErrContextWithError(errs.BadRequest, err))
	}

	if !req.isAllowanceTypeValid(allowanceType) {
		return c.JSON(http.StatusBadRequest, errs.NewErrContextWithError(errs.BadRequest, errors.New("allowance type is not valid: "+allowanceType)))
	}

	if !req.isAmountValid(allowanceType) {
		return c.JSON(http.StatusBadRequest, errs.NewErrContextWithError(errs.BadRequest, errors.New("amount is over limit")))
	}

	resp, err := h.svc.Process(c.Request().Context(), req, allowanceType)
	if err != nil {
		return c.JSON(errs.GenericError.HttpStatus, err)
	}

	return c.JSON(http.StatusOK, resp)
}
