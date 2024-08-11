package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Api struct {
}

type Hello struct {
	Hello string `json:"hello"`
}

func (a *Api) Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, Hello{Hello: "world"})
}
