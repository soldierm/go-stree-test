package components

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func RegisterMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
}
