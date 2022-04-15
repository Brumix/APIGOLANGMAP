package routes

import (
	"APIGOLANGMAP/controllers"

	"github.com/gin-gonic/gin"
)

// @Summary Realizar registro
// @Description Regista um utilizador
// @Accept  json
// @Produce  json
// @Router /auth/register [post]
// @Param evaluation body model.User true "Do register"
// @Success 201 {object} model.Claims
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
func RegisterUser(c *gin.Context) {
	controllers.RegisterHandler(c)
}

// @Summary Realizar autenticação
// @Description Autentica o utilizador e gera o token para os próximos acessos
// @Accept  json
// @Produce  json
// @Router /auth/login [post]
// @Param evaluation body model.User true "Do login"
// @Success 200 {object} model.Claims
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
func GenerateToken(c *gin.Context) {
	controllers.LoginHandler(c)
}

// @Summary Atualiza token de autenticação
// @Description Atualiza o token de autenticação do usuário invalidando o antigo
// @Accept  json
// @Produce  json
// @Router /auth/refresh_token [put]
// @param Authorization header string true "Token"
// @Success 200 {object} model.Claims
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
// @Failure 406 "Cannot invalidate old token"
func RefreshToken(c *gin.Context) {
	controllers.RefreshHandler(c)
}

// @Summary Realizar desautenticação
// @Description Desautentica o utilizador invalidando o token atual
// @Accept  json
// @Produce  json
// @Router /auth/logout [put]
// @Param evaluation body model.User true "Do logout"
// @param Authorization header string true "Token"
// @Success 200 "bool"
// @Failure 406 "Cannot log out"
func InvalidateToken(c *gin.Context) {
	controllers.LogoutHandler(c)
}

// @Summary Get User Info from token
// @Description retornas as informações de um dado user  atraves do seu token
// @Accept  json
// @Produce  json
// @Router /auth/getUser [get]
// @param Authorization header string true "Token"
// @Success 200 {object} model.User
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
func GetUserFromToken(c *gin.Context) {
	controllers.UserFromToken(c)
}
