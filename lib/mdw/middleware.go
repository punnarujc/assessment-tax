package mdw

import (
	"crypto/subtle"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/punnarujc/assessment-tax/lib/config"
)

func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(validateAdminCredential)
}

func validateAdminCredential(username, password string, c echo.Context) (bool, error) {
	adminCfg := config.New().GetAdminCredential()

	if subtle.ConstantTimeCompare([]byte(username), []byte(adminCfg.AdminUsername)) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte(adminCfg.AdminPassword)) == 1 {
		return true, nil
	}

	return false, nil
}
