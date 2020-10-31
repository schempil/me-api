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
	var input models.CreateCV

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cv := models.CV{FirstName: input.FirstName, LastName: input.LastName, Title: input.Title}
	models.DB.Create(&cv)

	c.JSON(http.StatusOK, gin.H{"data": cv})
}

func UpdateCV(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Sorry fo dat")
}

func DeleteCV(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Sorry fo dat")
}
