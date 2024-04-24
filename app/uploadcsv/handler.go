package uploadcsv

import (
	"net/http"

	"github.com/punnarujc/assessment-tax/lib/errs"
	"github.com/punnarujc/assessment-tax/lib/server"
)

type Handler interface {
	UploadCsv(c server.Context) error
}

type handlerImpl struct {
	svc Service
}

func NewHandler(svc Service) Handler {
	return &handlerImpl{
		svc: svc,
	}
}

func (h *handlerImpl) UploadCsv(c server.Context) error {
	taxFile, err := c.FormFile(PARAM_TAX_FILE)
	if err != nil {
		return c.JSON(errs.BadRequest.HttpStatus, errs.NewErrContextWithError(errs.BadRequest, err))
	}

	resp, err := h.svc.Process(c.Request().Context(), taxFile)
	if err != nil {
		return c.JSON(errs.GenericError.HttpStatus, errs.NewErrContextWithError(errs.GenericError, err))
	}

	return c.JSON(http.StatusOK, resp)
}
