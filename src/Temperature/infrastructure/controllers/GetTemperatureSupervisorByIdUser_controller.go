package controllers

import (
	"github.com/BryanChanona/backend_multi/src/Temperature/application/UseCase"
	"github.com/gin-gonic/gin"
)

type GetTemperatureSupervisorByIdUserController struct {
	useCase *UseCase.TemperatureSupervisorUc
}

func NewGetTemperatureSupervisorByIdUserController(useCase *UseCase.TemperatureSupervisorUc) *GetTemperatureSupervisorByIdUserController {
	return &GetTemperatureSupervisorByIdUserController{useCase: useCase}
}

func (controller *GetTemperatureSupervisorByIdUserController)Execute(ctx *gin.Context){
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
	
	temperatureSupervisors, err := controller.useCase.Execute(idUs)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": temperatureSupervisors})
}