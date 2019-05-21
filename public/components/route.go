package components

import (
	"github.com/labstack/echo"
	"html/template"
	"net/http"
)

var t = &IndexTemplate{
	templates: template.Must(template.ParseGlob("public/views/*.html")),
}

func RegisterRoute(e *echo.Echo) {
	e.Renderer = t
	e.GET("/", index)
}

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}
