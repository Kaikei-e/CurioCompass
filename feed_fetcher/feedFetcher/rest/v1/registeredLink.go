package v1

import "github.com/labstack/echo/v4"

func GetRegisteredLink(c echo.Context) error {
	return c.String(200, "https://echo.labstack.com")
}
