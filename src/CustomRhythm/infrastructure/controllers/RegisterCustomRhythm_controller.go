package controllers

import (
	"net/http"

	"github.com/BryanChanona/backend_multi/src/CustomRhythm/application/UseCase"
	"github.com/BryanChanona/backend_multi/src/CustomRhythm/domain"
	"github.com/gin-gonic/gin"
)

type RegisterCustomRhythmController struct {
	useCase *UseCase.RegisterCustomRhythmUC
}

func NewRegisterCustomRhythmController(useCase *UseCase.RegisterCustomRhythmUC) *RegisterCustomRhythmController {
	return &RegisterCustomRhythmController{useCase: useCase}
}

func (controller *RegisterCustomRhythmController) Execute(ctx *gin.Context){
	var customRhythm domain.CustomRhythmModel

	idUSer, exists := ctx.Get("id_user")
	if !exists {
		ctx.JSON(400, gin.H{"error": "id_user no encontrado"})
		return
	}
	if err := ctx.ShouldBindJSON(&customRhythm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}
	customRhythm.Id_user = idUSer.(int)

	if err := controller.useCase.Execute(customRhythm); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar el ritmo personalizado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Ritmo personalizado registrado correctamente"})




}