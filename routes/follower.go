package routes

import (
	"APIGOLANGMAP/controllers"

	"github.com/gin-gonic/gin"
)

// @Summary Obtem os Followers
// @Description Exibe a lista, sem todos os campos, de todos os followers
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Success 200 {array} model.Follower
// @Router /followers [get]
// @Failure 404 "Not found"
func GetAllFollowers(c *gin.Context) {
	controllers.GetAllFollowers(c)
}

// @Summary Associa um Follower(User) a um User
// @Description Associa um Follower a um User
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param follower body model.Follower	 true "Associate User as Follower"
// @Router /follower/assoc [post]
// @Success 200 {array} model.Follower
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
func AssociateFollower(c *gin.Context) {
	controllers.AssociateFollower(c)
}

// @Summary Desassocia um Follower(User) de um User
// @Description Desassocia um Follower de um User
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param follower body model.Follower true "Deassociate Follower from User"
// @Router /follower/deassoc [post]
// @Success 200 {array} model.Follower
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
func DeassociateFollower(c *gin.Context) {
	controllers.DeassociateFollower(c)
}
