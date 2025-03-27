package routes

import (
	"github.com/BryanChanona/backend_multi/src/Oxygen/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine){
	routes := router.Group("/oxygen")
	saveOxygenController := dependencies.GetSaveOxygenController().Execute
	getUserOxygenController := dependencies.GetOxygenationsController().Execute

	routes.POST("/",saveOxygenController)
	routes.GET("/",getUserOxygenController)

}