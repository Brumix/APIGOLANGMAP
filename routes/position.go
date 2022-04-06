package routes

import (
	"APIGOLANGMAP/controllers"
	"github.com/gin-gonic/gin"
)

//TODO SWAGGER
func RegisterLocation(c *gin.Context) {
	controllers.RegisterLocation(c)
}
