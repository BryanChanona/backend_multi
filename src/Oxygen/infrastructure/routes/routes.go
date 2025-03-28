package routes

import (
	"github.com/BryanChanona/backend_multi/src/Oxygen/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine){
	routes := router.Group("/oxygen")
	saveOxygenController := dependencies.GetSaveOxygenController().Execute
	getUserOxygenController := dependencies.GetOxygenationsController().Execute
	getUserOxygenByDate := dependencies.GetOxygenUserByDateController().Execute
	getUserOxygenById := dependencies.GetOxygenByIdController().Execute

	routes.POST("/",saveOxygenController)
	routes.GET("/",getUserOxygenController)
	routes.GET("/:idUser/:date",getUserOxygenByDate)
	routes.GET("/:idUser",getUserOxygenById)

}