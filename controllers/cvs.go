package controllers

import (
	"github.com/gin-gonic/gin"
	"models"
	"net/http"
)

func FindCVs(c *gin.Context) {
	var cvs []models.CV
	models.DB.Find(&cvs)

	c.JSON(http.StatusOK, gin.H{"data": cvs})
}
