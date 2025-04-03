package routes

import (
	"github.com/BryanChanona/backend_multi/src/CustomRhythm/infrastructure/dependencies"
	"github.com/BryanChanona/backend_multi/src/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine){
	routes := r.Group("customRhythm")
	registerCustomRhythm := dependencies.GetRegisterCustomRhythm().Execute
	updateCustomRhythm := dependencies.GetUpdateCustomRhythm().Execute
	routes.POST("/",middlewares.AuthMiddleware(),registerCustomRhythm)
	routes.PUT("/update",middlewares.AuthMiddleware(),updateCustomRhythm)
}