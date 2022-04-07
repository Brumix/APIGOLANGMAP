package controllers

import (
	"APIGOLANGMAP/model"
  "APIGOLANGMAP/services"
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

	}

	if errStore := repo.StorePosition(&position); errStore != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

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


func GetAllFollowers(c *gin.Context) {
	var evaluations []model.Evaluation

	services.Db.Find(&evaluations)

	if len(evaluations) <= 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Empty list!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": evaluations})
}

