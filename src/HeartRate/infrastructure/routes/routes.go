package routes

import (
	"github.com/BryanChanona/backend_multi/src/HeartRate/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine){
	routes := router.Group("/heartRate")
	getSaveHeartRateController := dependencies.GetHeartRateController().Execute
	getUserHeartRateController := dependencies.GetUserHeartRateController().Execute
	getUserHeartRateByDateController := dependencies.GetUserHeartRateByDateController().Execute	
	getUserHeartRateByIdController := dependencies.GetUserHeartRateByIdController().Execute

	routes.POST("/",getSaveHeartRateController)
	routes.GET("/",getUserHeartRateController)
	routes.GET("/:idUser/:date",getUserHeartRateByDateController)
	routes.GET("/:idUser",getUserHeartRateByIdController)


}