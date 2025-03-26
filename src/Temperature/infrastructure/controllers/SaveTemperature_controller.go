package controllers

import (
	"net/http"

	"github.com/BryanChanona/backend_multi/src/Temperature/application/UseCase"
	"github.com/BryanChanona/backend_multi/src/Temperature/domain"
	"github.com/gin-gonic/gin"
)

type SaveTemperatureController struct {
	UseCase *UseCase.SaveTemperatureUc
}


func NewSaveTemperatureController(uc *UseCase.SaveTemperatureUc) *SaveTemperatureController{
	return &SaveTemperatureController{UseCase: uc}
}

func (controller *SaveTemperatureController) Execute(ctx *gin.Context){
	var temperature domain.Temperature

	if err := ctx.ShouldBindJSON(&temperature); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Execute the use case to create the book
	err := controller.UseCase.Execute(temperature)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{"message": "Registered temperature"})
	}
}