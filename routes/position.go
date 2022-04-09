package routes

import (
	"APIGOLANGMAP/controllers"
	"APIGOLANGMAP/services"

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

// @Summary Exclui uma localização
// @Description Exclui uma localização selecionada
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param id path int true "Position ID"
// @Router /position/{id} [delete]
// @Success 200 "Delete succeeded!"
// @Failure 404 "None found!"
func DeleteLocation(c *gin.Context) {
	controllers.DeleteLocation(c)
}


// @Summary Obtem todas as localizações dos utilizadores com filtros
// @Description Exibe a lista de localizações dos utilizadores 
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Success 200 {array} model.Position
// @Router /position/filter [post]
// @Failure 404 "Location Not found"
// @Failure 400 "User Token Malformed"
func GetUsersLocationWithFilters(c *gin.Context){
	controllers.GetUsersLocationWithFilters(c)

// @Success 200 "Connection confirm"
// @Router /socket [get]
// @Failure 404 "Connection failed"
// @Failure 400 "User Token Malformed"
func WebSocket(c *gin.Context) {
	services.InitConnectionSocket(c)
  
}
