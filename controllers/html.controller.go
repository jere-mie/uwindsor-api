package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHTMLRoutes(rg *gin.RouterGroup) {
	rg.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Home",
		})
	})
	rg.GET("/guide", func(c *gin.Context) {
		c.HTML(http.StatusOK, "guide.html", gin.H{
			"title": "Guide",
		})
	})
	rg.GET("/sources", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sources.html", gin.H{
			"title": "Sources",
		})
	})
}
