package routes

import (
	"github.com/BryanChanona/backend_multi/src/Oxygen/infrastructure/dependencies"
	"github.com/BryanChanona/backend_multi/src/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine){
	routes := router.Group("/oxygen")
	saveOxygenController := dependencies.GetSaveOxygenController().Execute
	getUserOxygenController := dependencies.GetOxygenationsController().Execute
	getUserOxygenByDate := dependencies.GetOxygenUserByDateController().Execute
	getUserOxygenById := dependencies.GetOxygenByIdController().Execute
	getUserOxygenSupervisorById := dependencies.GetOxygenSupervisorByIdUserController().Execute

	routes.POST("/",saveOxygenController)
	routes.GET("/",getUserOxygenController)
	routes.GET("/:date",middlewares.AuthMiddleware(),getUserOxygenByDate)
	routes.GET("/oxygenById",middlewares.AuthMiddleware(),getUserOxygenById)
	routes.GET("/oxygenByIdSupervisor",middlewares.AuthSupervisorMiddleware(),getUserOxygenSupervisorById)

}