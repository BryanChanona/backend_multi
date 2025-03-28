package controller

import (
	"net/http"
	"strconv"

	"github.com/BryanChanona/backend_multi/src/Oxygen/application/UseCase"
	"github.com/gin-gonic/gin"
)

type OxygenByIdController struct {
	useCase *UseCase.OxygenByIdUc
}

func NewOxygenByIdController(useCase *UseCase.OxygenByIdUc)*OxygenByIdController{
	return &OxygenByIdController{useCase: useCase}
}

func (controller *OxygenByIdController)Execute(ctx *gin.Context){
	idUser, err := strconv.Atoi(ctx.Param("idUser"))

	if err != nil {
		// Manejo de error si el idUser no es válido
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	oxygenUsers, err := controller.useCase.Execute(idUser)
	if err != nil {
		// Manejo de error si hay un problema al obtener los datos
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": oxygenUsers})

}