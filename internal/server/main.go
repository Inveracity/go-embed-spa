package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/inveracity/go-embed-spa/ui"
)

func Server() http.Handler {
	r := gin.New()
	r.Use(cors.Default())
	r.StaticFS("/", ui.SvelteKitFS())

	return r
}

func ServerAPI() http.Handler {
	r := gin.New()
	r.Use(cors.Default())

	r.GET("/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	return r
}
