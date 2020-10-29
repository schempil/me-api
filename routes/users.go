package routes

import (
	"net/http"
	"person"

	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, person.GetPerson())
	})

	users.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "POST /users")
	})

	users.PUT("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "PUT /users")
	})

	users.DELETE("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "DELETE /users")
	})
}
