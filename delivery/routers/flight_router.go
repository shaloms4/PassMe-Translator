package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shaloms4/Pass-Me-Core-Functionality/delivery/controllers"
	Infrastructure "github.com/shaloms4/Pass-Me-Core-Functionality/infrastructure"
)

func SetupFlightRoutes(router *gin.Engine, controller *controllers.FlightController) {
	flights := router.Group("/flights")
	flights.Use(Infrastructure.AuthMiddleware())
	{
		flights.POST("", controller.CreateFlight)

		flights.GET("", controller.GetUserFlights)

		flights.GET("/:id", controller.GetFlightByID)

		flights.DELETE("/:id", controller.DeleteFlight)
	}
}
