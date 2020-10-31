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

func FindCV(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Sorry fo dat")
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
