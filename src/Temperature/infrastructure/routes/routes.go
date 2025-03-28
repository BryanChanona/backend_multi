package routes

import (
	"github.com/BryanChanona/backend_multi/src/Temperature/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/temperature")

	saveTemperatureController := dependencies.GetSaveTemperatureController().Execute
	getUSerTemperaturesController := dependencies.GetUserTemperaturesController().Execute
	getUserTemperatureByDateController := dependencies.GetUserTemperatureByDateController().Execute
	getUserTemperatureByIdController := dependencies.GetUserTemperaturesByIdController().Execute

	routes.POST("/", saveTemperatureController)
	routes.GET("/",getUSerTemperaturesController)
	routes.GET("/:idUser/:date",getUserTemperatureByDateController)
	routes.GET("/:idUser",getUserTemperatureByIdController)
}