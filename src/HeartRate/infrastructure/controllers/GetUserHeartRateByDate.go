package controllers

import (
	"net/http"
	"strconv"

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
	IdUser, err := strconv.Atoi(ctx.Param("idUser"))

	if err != nil {
		// Manejo de error si el idUser no es válido
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	date := ctx.Param("date")

	heartRateUser, err := controller.useCase.Execute(IdUser,date)
	if err != nil {
		// Manejo de error si hay un problema al obtener los datos
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": heartRateUser})
}