package controller

import (
	"net/http"

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
	date := ctx.Param("date")

	oxygenUser, err := controller.useCase.Execute(date,idUs)
	if err != nil {
		// Manejo de error si hay un problema al obtener los datos
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": oxygenUser})
}

