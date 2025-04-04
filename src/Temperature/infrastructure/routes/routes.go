package routes

import (
	"github.com/BryanChanona/backend_multi/src/Temperature/infrastructure/dependencies"
	"github.com/BryanChanona/backend_multi/src/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/temperature")

	saveTemperatureController := dependencies.GetSaveTemperatureController().Execute
	getUSerTemperaturesController := dependencies.GetUserTemperaturesController().Execute
	getUserTemperatureByDateController := dependencies.GetUserTemperatureByDateController().Execute
	getUserTemperatureByIdController := dependencies.GetUserTemperaturesByIdController().Execute
	getUserTemperatureSupervisorByIdController := dependencies.GetTemperatureSupervisorByIdUserController().Execute

	routes.POST("/", saveTemperatureController)
	routes.GET("/",getUSerTemperaturesController)
	routes.GET("/:date",middlewares.AuthMiddleware(),getUserTemperatureByDateController)
	routes.GET("/temperatureById",middlewares.AuthMiddleware(),getUserTemperatureByIdController)
	routes.GET("/temperatureByIdSupervisor",middlewares.AuthSupervisorMiddleware(),getUserTemperatureSupervisorByIdController)
}