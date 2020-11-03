package controllers

import (
	"github.com/gin-gonic/gin"
	"models"
	"net/http"
)

var (
	cvs []models.CV
	cv models.CV
)

func FindCVs(c *gin.Context) {
	models.DB.Find(&cvs)
	c.JSON(http.StatusOK, gin.H{"data": cvs})
}

func FindCV(c *gin.Context) {
	models.DB.First(&cvs, c.Param("id"))
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

	if err := models.DB.Where("id = ?", c.Param("id")).First(&cv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UpdateCV
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&cv).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": cv})

}

func DeleteCV(c *gin.Context) {

	if err := models.DB.Where("id = ?", c.Param("id")).First(&cv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&cv)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
