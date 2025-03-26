package routes

import (
	"github.com/BryanChanona/backend_multi/src/User/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/users")

	saveUserController := dependencies.GetSaveUserController().Execute
	routes.POST("/", saveUserController)
}