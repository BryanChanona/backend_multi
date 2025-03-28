package controller

import (
	"net/http"
	"strconv"

	"github.com/BryanChanona/backend_multi/src/Oxygen/application/UseCase"
	"github.com/gin-gonic/gin"
)

type OxygenByDateController struct {
	useCase *UseCase.GetOxygenByDateUc
}

func NewOxygenByDateController(useCase *UseCase.GetOxygenByDateUc)*OxygenByDateController{
	return &OxygenByDateController{useCase: useCase}
}

func (controller *OxygenByDateController)Execute(ctx *gin.Context){
	IdUser, err := strconv.Atoi(ctx.Param("idUser"))

	if err != nil {
		// Manejo de error si el idUser no es válido
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	date := ctx.Param("date")

	oxygenUser, err := controller.useCase.Execute(IdUser,date)
	if err != nil {
		// Manejo de error si hay un problema al obtener los datos
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": oxygenUser})
}

