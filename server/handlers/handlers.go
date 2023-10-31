package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"scoreboard-go/server/config"
	"scoreboard-go/server/models"
	"strconv"
)

func InsertData(c *gin.Context) {
	var data models.ScoreCard

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		return
	}
	exist, err := getRecordByName(data.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "failure",
			"error":  "Internal server",
		})
		return

	}
	if exist.Name != "" {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"status": "failure",
			"error":  "record already exist",
		})
		return

	}

	if data.SATScore > 30 {
		data.Passed = "Pass"
	} else {
		data.Passed = "Fail"

	}

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
	name := c.PostForm("name")
	newScore := c.PostForm("SATScore")

	data, err := getRecordByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "failure",
			"error":  "Internal server",
		})
		return

	}
	newScoreFloat, err := strconv.ParseFloat(newScore, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "failure",
			"error":  "Internal server",
		})
		return

	}
	data.SATScore = newScoreFloat

	if data.SATScore > 30 {
		data.Passed = "Pass"
	} else {
		data.Passed = "Fail"

	}
	config.DB.Save(&data)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func DeleteRecord(c *gin.Context) {

	name := c.PostForm("name")
	data, err := getRecordByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "failure",
			"error":  "record not found",
		})
		return

	}

	if data.Name == "" {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "failure",
			"error":  "record not found",
		})
		return

	}

	config.DB.Delete(&data)
	c.Status(http.StatusNoContent)
}
func GetRankStudent(c *gin.Context) {
	name := c.PostForm("name")

	result, err := getRecordByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "failure",
			"error":  "record not found",
		})
		return

	}
	//38
	var rank int64
	config.DB.Model(&models.ScoreCard{}).Where("sat_score > ?", result.SATScore).Count(&rank)
	fmt.Println("rank", rank)
	fmt.Println("sat_score", result.SATScore)

	c.JSON(http.StatusOK, gin.H{"name": name, "rank": rank + 1})
}
func getRecordByName(name string) (data models.ScoreCard, err error) {
	if err = config.DB.Where("name = ?", name).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil

		}

		return
	}
	return

}
func ViewSingleData(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failure",
			"error":  "name is required please give",
		})
		return
	}
	data, err := getRecordByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "failure",
			"error":  "Internal server",
		})
		return
	}
	if data.Name == "" {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "failure",
			"error":  "record not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"result": data,
	})

}
