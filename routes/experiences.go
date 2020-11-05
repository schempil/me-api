package routes

import (
	"controllers"
	"github.com/gin-gonic/gin"
)

func addExperiencesRoutes(rg *gin.RouterGroup) {
	experiences := rg.Group("/experiences")

	experiences.GET("/", controllers.FindExperiences)
	experiences.POST("/", controllers.CreateExperience)
	experiences.GET("/:id", controllers.FindExperience)
	experiences.PUT("/:id", controllers.UpdateExperience)
	experiences.DELETE("/:id", controllers.DeleteExperience)
}
