package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scoreboard-go/server/config"
	"scoreboard-go/server/models"
)

func InsertData(c *gin.Context) {
	var data models.ScoreCard
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data.Passed = data.SATScore > 30.0
	config.DB.Create(&data)
	c.JSON(http.StatusCreated, data)
}

func ViewAllData(c *gin.Context) {
	var data []models.ScoreCard
	config.DB.Find(&data)
	c.JSON(http.StatusOK, data)
}

func UpdateScore(c *gin.Context) {
	var data models.ScoreCard
	name := c.Param("name")

	if err := config.DB.Where("name = ?", name).First(&data).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	var newScore float64
	if err := c.ShouldBindJSON(&newScore); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data.SATScore = newScore
	data.Passed = data.SATScore > 30.0
	config.DB.Save(&data)
	c.Status(http.StatusNoContent)
}

func DeleteRecord(c *gin.Context) {
	var data models.ScoreCard
	name := c.Param("name")

	if err := config.DB.Where("name = ?", name).First(&data).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	config.DB.Delete(&data)
	c.Status(http.StatusNoContent)
}
func GetRankStudent(c *gin.Context) {
	name := c.Param("name")

	var result models.ScoreCard
	if err := config.DB.Where("name = ?", name).First(&result).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}

	var rank int64
	config.DB.Model(&models.ScoreCard{}).Where("SATScore > ?", result.SATScore).Count(&rank)

	c.JSON(200, gin.H{"name": name, "rank": rank + 1})
}
