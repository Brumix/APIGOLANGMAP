package routes

import (
	"APIGOLANGMAP/controllers"

	"github.com/gin-gonic/gin"
)

// @Summary Recupera as posições de todos os users
// @Description Exibe a lista, sem todos os campos, de todas as posições
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Success 200 {array} model.Position
// @Router /position [get]
// @Failure 404 "Not found"
func GetAllPositions(c *gin.Context) {
	controllers.GetAllPositions(c)
}

// @Summary Recupera uma posição pelo id do user
// @Description Exibe os detalhes da ultima posição do user através do seu id
// @ID get-position-by-int
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param id path int true "User ID"
// @Success 200 {object} model.Position
// @Router /position/{id} [get]
// @Failure 404 "Not found"
func GetPositionByUserId(c *gin.Context) {
	controllers.GetPositionByUserId(c)
}

// @Summary Adicionar uma posição
// @Description Cria uma posição sobre um utilizador
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param position body model.Position true "Add position"
// @Router /position [post]
// @Success 201 {object} model.Position
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
func AddPosition(c *gin.Context) {
	controllers.AddPosition(c)
}
