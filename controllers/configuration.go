package controllers

import (
	"APIGOLANGMAP/model"
	"APIGOLANGMAP/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateAlertTime(c *gin.Context) {
	var user model.User

	id := c.Param("id")
	services.Db.First(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "User not found!"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Check request!"})
		return
	}

	result := services.Db.Save(user)

	if result.RowsAffected != 0 {
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusOK, "message": "Success!", "User ID": user.ID})
		return
	}

	c.JSON(http.StatusNotAcceptable, gin.H{"status": http.StatusNotAcceptable, "message": "Cannot be Updated!"})
}
