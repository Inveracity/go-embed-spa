package server

import (
	"net/http"

	"github.com/inveracity/go-embed-spa/internal/common"
	"github.com/labstack/echo/v4"
)

type Api struct {
}

func (a *Api) Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, common.Hello{Hello: "world"})
}
