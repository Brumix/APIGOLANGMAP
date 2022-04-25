package controllers

import (
	"APIGOLANGMAP/model"
	"APIGOLANGMAP/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ActivateSOS(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}
	services.Db.Find(&user, "username = ?", user.Username)

	activated := user.IsSOSActivated

	if activated == true {
		c.JSON(http.StatusNotAcceptable, gin.H{"status": http.StatusNotAcceptable, "message": "SOS is already activated!"})
	}

	if activated == false {
		activated = true
	}

	user.IsSOSActivated = activated
	services.Db.Save(&user)
}

func DesactivateSOS(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}
	services.Db.Find(&user, "username = ?", user.Username)

	activated := user.IsSOSActivated

	if activated == false {
		c.JSON(http.StatusNotAcceptable, gin.H{"status": http.StatusNotAcceptable, "message": "SOS is already desactivated!"})
	}

	if activated == true {
		activated = false
	}

	user.IsSOSActivated = activated
	services.Db.Save(&user)
}

func GetAllUsers(c *gin.Context) {
	var users []model.User
	services.Db.Find(&users)
	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}
