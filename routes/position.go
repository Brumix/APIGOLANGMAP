package routes

import (
	"APIGOLANGMAP/controllers"
	"github.com/gin-gonic/gin"
)

// @Summary Adicionar uma localizaçao
// @Description Cria uma localizacao de um utilizador em especifico
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param evaluation body model.Position true "Add Location"
// @Router /position [post]
// @Success 201 {object} model.Position
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
func RegisterLocation(c *gin.Context) {
	controllers.RegisterLocation(c)
}

// @Summary Exclui uma localização
// @Description Exclui uma localização selecionada
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param id path int true "Position ID"
// @Router /position/{id} [delete]
// @Success 200 {object} model.Position
// @Failure 404 "Not found"
func DeleteLocation(c *gin.Context) {
	controllers.DeleteLocation(c)
}
