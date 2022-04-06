package main

import (
	"APIGOLANGMAP/model"
	"APIGOLANGMAP/routes"
	"APIGOLANGMAP/services"

	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/postgres"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var identityKey = "id"

const UserAccess  = false
const AdminAccess = true

func init() {
	services.OpenDatabase()
	services.Db.AutoMigrate(&model.Evaluation{})
	services.Db.AutoMigrate(&model.User{})
	services.Db.AutoMigrate(&model.RevokedToken{})
	services.CreateAdmin()
}

func main() {

	services.FormatSwagger()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// NO AUTH
	router.GET("/api/v1/echo", routes.EchoRepeat)

	// AUTH
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	evaluation := router.Group("/api/v1/evaluation")
	evaluation.Use(services.AuthorizationRequired(AdminAccess)) // admin just to test who can access this route
	{
		evaluation.POST("/", routes.AddEvaluation)
		evaluation.GET("/", routes.GetAllEvaluation)
		evaluation.GET("/:id", routes.GetEvaluationById)
		evaluation.PUT("/:id", routes.UpdateEvaluation)
		evaluation.DELETE("/:id", routes.DeleteEvaluation)
	}

	/* localization := router.Group("/api/v1/localization")
	localization.Use(services.AuthorizationRequired(AdminAccess))
	{
		localization.GET("/", routes.GetAllLocalizations)
		localization.GET("/:id", routes.GetLocalizationById)
		localization.GET("/:id", routes.GetLocalizationUsersByFilter)
	} */

	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/login", routes.GenerateToken)
		auth.PUT("/logout", routes.InvalidateToken)
		auth.POST("/register", routes.RegisterUser)
		auth.PUT("/refresh_token", services.AuthorizationRequired(UserAccess), routes.RefreshToken)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
