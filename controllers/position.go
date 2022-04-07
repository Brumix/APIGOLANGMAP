package controllers

import (
	"APIGOLANGMAP/model"
	"APIGOLANGMAP/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

var repo = repository.NewCrudPositions()

func RegisterLocation(c *gin.Context) {
	var position model.Position

	err := c.Bind(&position)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return

func GetAllPositions(c *gin.Context) {
	var positions []model.Position
	}

	services.Db.Find(&positions)
	if errStore := repo.StorePosition(&position); errStore != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if len(positions) <= 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Empty list!"})
	c.JSON(http.StatusOK, gin.H{"msg": "Position register with success!!",
		"Position": position})
	return
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
