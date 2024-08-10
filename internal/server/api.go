package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Api struct {
}

func (a *Api) Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"hello": "World!"})
}
