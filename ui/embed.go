package ui

import (
	"embed"

	"github.com/labstack/echo/v4"
)

//go:generate npm run build
//go:embed all:dist
var distDir embed.FS

// DistDirFS contains the embedded dist directory files (without the "dist" prefix)
var DistDirFS = echo.MustSubFS(distDir, "dist")
