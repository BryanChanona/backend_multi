package routes

import (
	"github.com/BryanChanona/backend_multi/src/Temperature/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/temperature")

	saveTemperatureController := dependencies.GetSaveTemperatureController().Execute
	getUSerTemperaturesController := dependencies.GetUserTemperaturesController().Execute

	routes.POST("/", saveTemperatureController)
	routes.GET("/",getUSerTemperaturesController)
}