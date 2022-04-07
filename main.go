package main

import (
	"APIGOLANGMAP/model"
	"APIGOLANGMAP/repository"
	"APIGOLANGMAP/routes"
	"APIGOLANGMAP/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gorm.io/driver/postgres"
)

var identityKey = "id"

const UserAccess = false
const AdminAccess = true

func init() {
	services.OpenDatabase()
	services.Db.AutoMigrate(&model.Evaluation{})
	services.Db.AutoMigrate(&model.User{})
	services.Db.AutoMigrate(&model.RevokedToken{})
	services.Db.AutoMigrate(&model.Position{})
	services.Db.Exec("alter table positions add column geolocation geography(point)")
	services.Db.AutoMigrate(&model.Follower{})
	services.CreateAdmin()
	//	services.CloseDatabase()
	repository.GetDataBase(services.Db)
	services.StartService()
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

	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/login", routes.GenerateToken)
		auth.POST("/logout", services.AuthorizationRequired(UserAccess), routes.InvalidateToken)
		auth.POST("/register", routes.RegisterUser)
		auth.PUT("/refresh_token", services.AuthorizationRequired(UserAccess), routes.RefreshToken)
	}

	position := router.Group("/api/v1/position")
	{
		position.POST("/", routes.RegisterLocation)
		position.DELETE("/:id", routes.DeleteLocation)

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
