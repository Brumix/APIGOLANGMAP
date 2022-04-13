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

func init() {
	services.OpenDatabase()
	services.Db.AutoMigrate(&model.User{})
	services.Db.AutoMigrate(&model.RevokedToken{})
	services.Db.AutoMigrate(&model.Position{})
	services.Db.Exec("alter table positions add column if not exists geolocation geography(point)")
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

	// AUTH
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	follower := router.Group("/api/v1/follower")
	follower.Use(services.AuthorizationRequired())
	{
		follower.GET("/", routes.GetAllFollowers)
		follower.POST("/assoc", routes.AssociateFollower)
		follower.POST("/deassoc", routes.DeassociateFollower)
	}

	alertTime := router.Group("/api/v1/alert")
	alertTime.Use(services.AuthorizationRequired())
	{
		alertTime.PUT("/time/", routes.UpdateAlertTime)

	}

	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/login", routes.GenerateToken)
		auth.POST("/logout", services.AuthorizationRequired(), routes.InvalidateToken)
		auth.POST("/register", routes.RegisterUser)
		auth.PUT("/refresh_token", services.AuthorizationRequired(), routes.RefreshToken)
	}

	position := router.Group("/api/v1/position")
	position.Use(services.AuthorizationRequired())
	{
		position.POST("/", routes.RegisterLocation)
		position.GET("/", routes.GetMyLocation)
		position.POST("/history", routes.GetLocationHistory)

		position.DELETE("/", routes.DeleteLocation)
		position.POST("/filter", routes.GetUsersLocationWithFilters)
		position.POST("/users_under_xkms", routes.GetAllUsersUnderXKms)

	}

	sos := router.Group("/api/v1/sos")
	position.Use(services.AuthorizationRequired())
	{
		sos.POST("/activate", routes.ActivateSOS)
		sos.POST("/desactivate", routes.DesactivateSOS)
	}

	router.GET("/socket", routes.WebSocket)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")

}
