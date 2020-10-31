package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func addCVRoutes(rg *gin.RouterGroup) {
	cvs := rg.Group("/cvs")

	cvs.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "GET /cvs")
	})

	cvs.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "POST /cvs")
	})

	cvs.GET("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, "GET /cvs")
	})

	cvs.PUT("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, "PUT /cvs")
	})

	cvs.DELETE("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, "DELETE /cvs")
	})
}
