package routes

import (
	"github.com/gin-gonic/gin"
	"models"
)

var (
	router = gin.Default()
)

func Run() {
	getRoutes()
	models.ConnectDataBase()
	router.Run(":8080")
}

func getRoutes() {
	v1 := router.Group("/v1")
	addCVRoutes(v1)
}
