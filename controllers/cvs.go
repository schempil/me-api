package controllers

import (
	"github.com/gin-gonic/gin"
	"models"
	"net/http"
)

var (
	cvs []models.CV
)

func FindCVs(c *gin.Context) {
	models.DB.Find(&cvs)
	c.JSON(http.StatusOK, gin.H{"data": cvs})
}

func FindCV(c *gin.Context) {
	id := c.Param("id")
	models.DB.First(&cvs, id)
	c.JSON(http.StatusOK, gin.H{"data": cvs})
}

func CreateCV(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Sorry fo dat")
}

func UpdateCV(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Sorry fo dat")
}

func DeleteCV(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Sorry fo dat")
}
