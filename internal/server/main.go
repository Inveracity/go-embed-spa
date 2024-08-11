package server

import (
	"fmt"

	"github.com/inveracity/go-embed-spa/internal/server/events"
	"github.com/inveracity/go-embed-spa/ui"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Server(bind string, port uint) {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.ListenerNetwork = "tcp4"

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${status} ${uri} ${remote_ip}\n",
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

	fmt.Printf("\n   running http://%s:%v\n\n", bind, port)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%v", bind, port)))
}
