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

  
// @Summary Obtem os Followers
// @Description Exibe a lista, sem todos os campos, de todas as avaliações
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Success 200 {array} model.Evaluation
// @Router /evaluation [get]
// @Failure 404 "Not found"
func GetAllFollowers(c *gin.Context) {
	controllers.GetAllEvaluations(c)
}
