package controllers

import (
	"APIGOLANGMAP/model"
	"APIGOLANGMAP/repository"
	"APIGOLANGMAP/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Location struct {
	gorm.Model `swaggerignore:"true" json:"-"`
	Start      string `json:"start" binding:"required,min=8,max=10"`
	End        string `json:"end" binding:"required,min=8,max=10"`
	Limit      int    `json:"limit"`
}

var repo = repository.NewCrudPositions()

func RegisterLocation(c *gin.Context) {
	var position model.Position
	userID, errAuth := c.Get("userid")
	if errAuth == false {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User Auth Token Malformed!"})
		return
	}
	err := c.Bind(&position)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return

	}
	position.UserID = userID.(uint)
	fmt.Println(">>>", position)
	if errStore := repo.StorePosition(&position); errStore != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Position register with success!!",
		"Position": position})
	return
}

func GetLastLocation(c *gin.Context) {
	var position model.Position
	userID, errAuth := c.Get("userid")
	if errAuth == false {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User Auth Token Malformed!"})
		return
	}

	if err := services.Db.Where("user_id = ?", userID).Order("created_at DESC").First(&position).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User ID Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "Got My Current Location", "location": position})
	return
}

func GetLocationHistory(c *gin.Context) {
	var location Location
	var positions []model.Position
	userID, errAuth := c.Get("userid")
	if errAuth == false {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User Auth Token Malformed!"})
		return
	}

	if err := c.ShouldBindJSON(&location); err != nil {
		// If Body Comes Empty -> Get Last Location
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Check Syntax!"})
		return
	}
	var defaultLimit = 2
	if location.Limit != 0 { // Verify If Body Has 'limit' Field
		defaultLimit = location.Limit
	}
	fmt.Println(">>> ", defaultLimit)

	if err := services.Db.Where("user_id = ?", userID).Order("created_at DESC").Find(&positions).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User ID Not Found"})
		return
	}

	if location.Start != "" && location.End != "" && location.Limit == 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "My Locations History", "locations": positions})
		return
	}

	// Has Body
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "My Locations History By Filter"})
	return

	/*
		if err := services.Db.Where("user_id = ?", userID).Order("created_at DESC").Find(&positions).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User ID Not Found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "My Locations History", "locations": positions})*/
}

func DeleteLocation(c *gin.Context) {
	var position model.Position
	err := c.Bind(&position)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return

	}
	if errDelete := repo.DeletePosition(&position); errDelete != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Position deleted with success!!",
		"Position": position})
	return
}
