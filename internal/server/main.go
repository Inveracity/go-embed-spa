package server

import (
	"fmt"
	"net/http"

	"github.com/inveracity/go-embed-spa/ui"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Server(port uint) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api", hello)
	e.GET("/*", echo.StaticDirectoryHandler(ui.DistDirFS, false))

	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%v", port)))
}

// Handler
func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"Hello": "World!"})
}
