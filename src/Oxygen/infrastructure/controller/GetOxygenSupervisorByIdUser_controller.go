package controller

import (
	"github.com/BryanChanona/backend_multi/src/Oxygen/application/UseCase"
	"github.com/gin-gonic/gin"
)

type OxygenSupervisorController struct {
	useCase *UseCase.OxygenSupervisorUc
}

func NewOxygenSupervisorController(useCase *UseCase.OxygenSupervisorUc) *OxygenSupervisorController {
	return &OxygenSupervisorController{useCase: useCase}
}

func (controller *OxygenSupervisorController)Execute(ctx *gin.Context) {
	idUser, exists := ctx.Get("id_user")
	if !exists {
		ctx.JSON(400, gin.H{"error": "ID de usuario no proporcionado"})
		return
	}

	idUs, ok := idUser.(int)
	if !ok {
		ctx.JSON(400, gin.H{"error": "ID de usuario no válido"})
		return
	}

	oxygenSupervisors, err := controller.useCase.Execute(idUs)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": oxygenSupervisors})
}
