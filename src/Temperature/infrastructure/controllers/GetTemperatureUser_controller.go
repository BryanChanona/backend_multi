package controllers

import (
	"net/http"

	"github.com/BryanChanona/backend_multi/src/Temperature/application/UseCase"
	"github.com/gin-gonic/gin"
)

type GetTemperatureUsersController struct {
	useCase *UseCase.GetTemperatureUser
}

func NewGetTemperatureUsersController(uc *UseCase.GetTemperatureUser) *GetTemperatureUsersController{
	return &GetTemperatureUsersController{useCase: uc}
}

func (controller *GetTemperatureUsersController)Execute(ctx *gin.Context){
	userTemperatures,err := controller.useCase.Execute()

	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"User temperatures": userTemperatures})
}