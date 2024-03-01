package rest

import (
	v1 "feedFetcher/rest/v1"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitServer() {
	e := echo.New()
	e.Use(
		middleware.SecureWithConfig(
			middleware.SecureConfig{
				XSSProtection:      "1; mode=block",
				ContentTypeNosniff: "nosniff",
				XFrameOptions:      "SAMEORIGIN",
			},
		),
	)

	e.GET("/system/health", healthCheck)
	v1Group := e.Group("/api/v1")
	v1Group.Use()
	{
		v1Group.GET("/registeredLink", v1.GetRegisteredLink)
	}

	e.Logger.Fatal(e.Start(":9000"))
}

func healthCheck(c echo.Context) error {
	return c.String(200, "OK")
}
