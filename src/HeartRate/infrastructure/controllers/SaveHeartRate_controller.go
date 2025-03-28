package controllers

import (
	"net/http"

	"github.com/BryanChanona/backend_multi/src/HeartRate/application/UseCase"
	"github.com/BryanChanona/backend_multi/src/HeartRate/domain"
	"github.com/gin-gonic/gin"
)

type SaveHeartRateController struct {
	useCase *UseCase.SaveHeartRateUc
}

func NewSaveHeartRateController(useCase *UseCase.SaveHeartRateUc)*SaveHeartRateController{
	return &SaveHeartRateController{useCase: useCase}
}

func (controller *SaveHeartRateController) Execute(ctx *gin.Context){
	var heartRate domain.HeartRate

	if err := ctx.ShouldBindJSON(&heartRate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := controller.useCase.Execute(heartRate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{"message": "Registered Heart Rate"})
	}

}