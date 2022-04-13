package controllers

import (
	"APIGOLANGMAP/model"
	"APIGOLANGMAP/services"
	"github.com/gin-gonic/gin"
	"net/http"
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
