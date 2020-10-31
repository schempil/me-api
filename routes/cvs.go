package routes

import (
	"controllers"
	"github.com/gin-gonic/gin"
)

func addCVRoutes(rg *gin.RouterGroup) {
	cvs := rg.Group("/cvs")

	cvs.GET("/", controllers.FindCVs)
	cvs.POST("/", controllers.CreateCV)
	cvs.GET("/:id", controllers.FindCV)
	cvs.PUT("/:id", controllers.UpdateCV)
	cvs.DELETE("/:id", controllers.DeleteCV)
}
