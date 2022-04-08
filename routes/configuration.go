package routes

import (
	"APIGOLANGMAP/controllers"

	"github.com/gin-gonic/gin"
)

// @Summary Atualiza a periodicidade de alerta
// @Description Atualiza a periodicidade de alerta determinando o tempo máximo até dar uma pessoa como perdida
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param evaluation body model.User true "Udpdate Alert"
// @Param id path int true "User ID"
// @Router /alert/time/ [put]
// @Success 200 {object} model.User
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
// @Failure 406 "Not acceptable"
// @Failure 401 "Unauthorized"
func UpdateAlertTime(c *gin.Context) {
	controllers.UpdateAlertTime(c)
}
