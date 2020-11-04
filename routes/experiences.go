package routes

import (
	"controllers"
	"github.com/gin-gonic/gin"
)

func addExperiencesRoutes(rg *gin.RouterGroup) {
	experiences := rg.Group("/experiences")

	experiences.GET("/", controllers.FindCVs)
	experiences.POST("/", controllers.CreateCV)
	experiences.GET("/:id", controllers.FindCV)
	experiences.PUT("/:id", controllers.UpdateCV)
	experiences.DELETE("/:id", controllers.DeleteCV)
}
