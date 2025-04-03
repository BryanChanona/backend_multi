package controllers

import (
	"net/http"
	"github.com/BryanChanona/backend_multi/src/Temperature/application/UseCase"
	"github.com/gin-gonic/gin"
)

type TemperatureByIdController struct {
	useCase *UseCase.TemperatureByIdUc
}
func NewTemperatureByIdController(useCase *UseCase.TemperatureByIdUc)*TemperatureByIdController{
	return &TemperatureByIdController{useCase: useCase}
}

func (controller *TemperatureByIdController)Execute(ctx *gin.Context){
	idUser, exists := ctx.Get("id_user")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario no proporcionado"})
		return
	}

	idUs, ok := idUser.(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario no válido"})
		return
	}
	userTemperatures, err := controller.useCase.Execute(idUs)
	if err != nil {
		// Manejo de error si hay un problema al obtener los datos
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": userTemperatures})

}