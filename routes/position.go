package routes

import (
	"APIGOLANGMAP/controllers"
	"github.com/gin-gonic/gin"
)

//TODO SWAGGER
func RegisterLocation(c *gin.Context) {
	controllers.RegisterLocation(c)
}
func DeleteLocation(c *gin.Context) {
	controllers.DeleteLocation(c)
}

//TODO Swagger
func GetAllUsersUnderXKms(c *gin.Context) {
	controllers.GetAllUsersUnderXKms(c)
}
