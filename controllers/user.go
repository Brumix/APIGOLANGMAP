package controllers

import (
	"APIGOLANGMAP/model"
	"APIGOLANGMAP/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ActivateSOS(c *gin.Context) {
	var user model.User

	services.OpenDatabase()

	if err := services.Db.First(&user, "username = ?", c.GetString("username")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User doesnt exists!"})
		return
	}

	activated := user.IsSOSActivated

	if activated == true {
		c.JSON(http.StatusNotAcceptable, gin.H{"status": http.StatusNotAcceptable, "message": "SOS is already activated!"})
	}

	if activated == false {
		activated = true
	}

	user.IsSOSActivated = activated
	services.Db.Save(&user)
	services.CloseDatabase()
}

func DesactivateSOS(c *gin.Context) {
	var user model.User

	services.OpenDatabase()

	if err := services.Db.First(&user, "username = ?", c.GetString("username")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User doesnt exists!"})
		return
	}

	activated := user.IsSOSActivated

	if activated == false {
		c.JSON(http.StatusNotAcceptable, gin.H{"status": http.StatusNotAcceptable, "message": "SOS is already desactivated!"})
	}

	if activated == true {
		activated = false
	}

	user.IsSOSActivated = activated
	services.Db.Save(&user)
	services.CloseDatabase()
}

