package server

import (
	"github.com/labstack/echo/v4"
	"github.com/punnarujc/assessment-tax/lib/errs"
)

type Context struct {
	echo.Context
}

func NewContext(c echo.Context) *Context {
	return &Context{
		Context: c,
	}
}

func (c *Context) JSON(code int, obj interface{}) error {
	switch obj.(type) {

	case errs.ErrContext:
		return c.Context.JSON(code, obj)
	}

	return c.Context.JSON(code, obj)
}
