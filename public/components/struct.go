package components

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
)

// 实现 Renderer 接口
type IndexTemplate struct {
	templates *template.Template
}

func (t *IndexTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
