package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shaloms4/Pass-Me-Core-Functionality/delivery/controllers"
	Infrastructure "github.com/shaloms4/Pass-Me-Core-Functionality/infrastructure"
)

func SetupUserRoutes(router *gin.Engine, controller *controllers.UserController) {
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)

	auth := router.Group("/profile")
	auth.Use(Infrastructure.AuthMiddleware())
	{
		auth.GET("/", controller.GetProfile)
		auth.PUT("/username", controller.ChangeUsername)
		auth.PUT("/password", controller.ChangePassword)
	}
}
