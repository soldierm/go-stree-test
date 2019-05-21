package components

import "github.com/labstack/echo"

func RegisterStatic(e *echo.Echo) {
	e.Static("/assets", "public/assets")
}
