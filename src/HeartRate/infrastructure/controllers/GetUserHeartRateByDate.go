package controllers

import (
	"net/http"

	"github.com/BryanChanona/backend_multi/src/HeartRate/application/UseCase"
	"github.com/gin-gonic/gin"
)

type UserHeartRateByDateController struct {
	useCase *UseCase.HeartRateByDateUc
}

func NewUserHeartRateByDateController(useCase *UseCase.HeartRateByDateUc) *UserHeartRateByDateController{
	return &UserHeartRateByDateController{useCase: useCase}
}

func (controller *UserHeartRateByDateController)Execute(ctx *gin.Context){
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

	heartRateUser, err := controller.useCase.Execute(idUs,date)
	if err != nil {
		// Manejo de error si hay un problema al obtener los datos
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": heartRateUser})
}