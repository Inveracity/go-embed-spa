package server

import (
	"fmt"

	"github.com/inveracity/go-embed-spa/ui"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Server(port uint) {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${status} ${uri}\n",
	}))
	e.Use(middleware.Recover())

	api := Api{}
	e.GET("/api", api.Hello)
	e.GET("/*", echo.StaticDirectoryHandler(ui.DistDirFS, false))

	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%v", port)))
}
