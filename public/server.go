package main

import (
	"github.com/labstack/echo"
	"go-stress-test/public/components"
	"os"
)

func main() {
	err := os.Setenv("CURRENT_ENV", "web-server")
	if err != nil {
		return
	}
	e := echo.New()
	components.RegisterMiddleware(e)
	components.RegisterStatic(e)
	components.RegisterRoute(e)
	e.Logger.Fatal(e.Start(":1323"))
}
