package controllers

import (
	"net/http"

	"github.com/BryanChanona/backend_multi/src/CustomRhythm/application/UseCase"
	"github.com/BryanChanona/backend_multi/src/CustomRhythm/domain"
	"github.com/gin-gonic/gin"
)

type UpdateCustomRhythmController struct {
	useCase *UseCase.UpdateCustomRhythmUC
}

func NewUpdateCustomRhythmController(useCase *UseCase.UpdateCustomRhythmUC) *UpdateCustomRhythmController {
	return &UpdateCustomRhythmController{useCase: useCase}
}

func (controller *UpdateCustomRhythmController)Execute(ctx *gin.Context){
	var customRhythm domain.CustomRhythmModel
	idUser, exists :=ctx.Get("id_user")
	if !exists {
		ctx.JSON(400, gin.H{"error": "User ID not found "})
		return
	}

	if err := ctx.ShouldBindJSON(&customRhythm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	err := controller.useCase.UpdateCustomRhythm(idUser.(int), customRhythm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el ritmo personalizado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Ritmo personalizado actualizado correctamente"})



	
}