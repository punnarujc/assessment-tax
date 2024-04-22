package server

import "github.com/labstack/echo/v4"

func Convert(handlerFunc HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return handlerFunc(Context{
			Context: c,
		})
	}
}
