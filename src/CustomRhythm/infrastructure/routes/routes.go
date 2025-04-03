package routes

import (
	"github.com/BryanChanona/backend_multi/src/CustomRhythm/infrastructure/dependencies"
	"github.com/BryanChanona/backend_multi/src/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine){
	routes := r.Group("customRhythm")
	registerCustomRhythm := dependencies.GetRegisterCustomRhythm().Execute
	routes.POST("/",middlewares.AuthMiddleware(),registerCustomRhythm)
}