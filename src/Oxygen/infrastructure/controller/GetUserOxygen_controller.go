package controller

import (
	"net/http"

	"github.com/BryanChanona/backend_multi/src/Oxygen/application/UseCase"
	"github.com/gin-gonic/gin"
)

type UserOxygenController struct {
	useCase *UseCase.GetUserOxygenUC
}

func NewUserOxygenController(uc *UseCase.GetUserOxygenUC) *UserOxygenController{
	return &UserOxygenController{useCase: uc}
}

func (controller *UserOxygenController)Execute(ctx *gin.Context){
	oxygenations,err := controller.useCase.Execute()

	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"User oxygenations": oxygenations})
}
