package controllers

import (
	"APIGOLANGMAP/model"
	"APIGOLANGMAP/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPositions(c *gin.Context) {
	var positions []model.Position

	services.Db.Find(&positions)

	if len(positions) <= 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Empty list!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": positions})
}

func GetPositionByUserId(c *gin.Context) {
	var position model.Position
	user_id := c.Param("user_id")

	services.Db.First(&position, user_id)
	if position.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": position})
}

func AddPosition(c *gin.Context) {
	var position model.Position

	var userid, err = c.Get("userid")

	if err == false {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User Auth Token Malformed!"})
		return
	}

	if err := c.ShouldBindJSON(&position); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Check syntax!"})
		return
	}

	position.UserID = userid.(uint)

	services.Db.Save(&position)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Create successful!", "resourceId": position.ID})
}
