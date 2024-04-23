package server

import (
	"github.com/labstack/echo/v4"
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
	// switch obj.(type) {

	// case errs.ErrContext:
	// 	return c.Context.JSON(code, obj)
	// }

	return c.Context.JSON(code, obj)
}

func (c *Context) Bind(obj interface{}) error {
	err := c.Context.Bind(obj)
	if err != nil {
		return err
	}

	return nil
}

func (c *Context) Param(name string) string {
	return c.Context.Param(name)
}
