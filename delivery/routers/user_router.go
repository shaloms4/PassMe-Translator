package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shaloms4/Pass-Me-Core-Functionality/delivery/controllers"
)

func SetupUserRoutes(router *gin.Engine, controller *controllers.UserController) {
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)
}
