package server

import (
	"fmt"

	"github.com/inveracity/go-embed-spa/internal/server/events"
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	api := Api{}
	e.GET("/api", api.Hello)
	e.GET("/stream/time", events.TimeStream)
	e.GET("/stream/memory", events.StreamMemory)
	e.GET("/stream/syslog", events.StreamSyslog)
	e.GET("/*", echo.StaticDirectoryHandler(ui.DistDirFS, false))

	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%v", port)))
}
