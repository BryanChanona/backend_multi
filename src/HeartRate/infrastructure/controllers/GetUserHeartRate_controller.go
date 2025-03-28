package controllers

import (
	"net/http"

	"github.com/BryanChanona/backend_multi/src/HeartRate/application/UseCase"
	"github.com/gin-gonic/gin"
)

type UserHeartRateController struct {
	useCase *UseCase.UserHeartRateUc
}

func NewUserHeartRateController(useCase *UseCase.UserHeartRateUc) *UserHeartRateController{
	return &UserHeartRateController{useCase: useCase}
}

func (controller *UserHeartRateController) Execute(ctx *gin.Context){
	 userHeartRates, err := controller.useCase.Execute()

	 
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"User Heart Rates": userHeartRates})





}