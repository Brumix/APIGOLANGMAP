package routes

import (
	"APIGOLANGMAP/controllers"

	"github.com/gin-gonic/gin"
)

//TODO SWAGGER
func RegisterLocation(c *gin.Context) {
	controllers.RegisterLocation(c)
}

// @Summary Obter a última localização do utilizador
// @Description Exibe a lista da última localização do utilizador
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Success 200 {object} model.Position
// @Router /position [get]
// @Failure 404 "User Not found"
// @Failure 400 "User Token Malformed"
func GetMyLocation(c *gin.Context) {
	controllers.GetLastLocation(c)
}

// @Summary Obtem todas as localizações do utilizador
// @Description Exibe a lista de todas as localizações do utilizador
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Success 200 {array} model.Position
// @Router /position/history [get]
// @Failure 404 "User Not found"
// @Failure 400 "User Token Malformed"
func GetLocationHistory(c *gin.Context) {
	controllers.GetLocationHistory(c)
}

func DeleteLocation(c *gin.Context) {
	controllers.DeleteLocation(c)
}
