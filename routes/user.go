package routes

import (
	"APIGOLANGMAP/controllers"
	"github.com/gin-gonic/gin"
)

//TODO Swagger
func ActivateSOS(c *gin.Context) {
	controllers.ActivateSOS(c)
}

//TODO Swagger
func DesactivateSOS(c *gin.Context) {
	controllers.DesactivateSOS(c)
}
