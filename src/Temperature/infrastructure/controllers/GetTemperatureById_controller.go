package controllers

import (
	"net/http"
	"strconv"

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
	idUser , err := strconv.Atoi(ctx.Param("idUser"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	userTemperatures, err := controller.useCase.Execute(idUser)
	if err != nil {
		// Manejo de error si hay un problema al obtener los datos
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": userTemperatures})

}