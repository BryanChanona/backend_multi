package controller

import (
	"net/http"

	"github.com/BryanChanona/backend_multi/src/Oxygen/application/UseCase"
	"github.com/BryanChanona/backend_multi/src/Oxygen/domain"
	"github.com/gin-gonic/gin"
)

type SaveOxygenController struct {
	useCase *UseCase.SaveOxygenUc
}

func NewSaveOxygenController(uc *UseCase.SaveOxygenUc) *SaveOxygenController{
	return &SaveOxygenController{useCase: uc}
}

func (controller *SaveOxygenController) Execute(ctx *gin.Context){
	var oxygen domain.OxygenModel

	if err := ctx.ShouldBindJSON(&oxygen); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := controller.useCase.Execute(oxygen)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{"message": "Registered oxygen"})
	}

}