package controllers

import (
	"net/http"

	"github.com/BryanChanona/backend_multi/src/Temperature/application/UseCase"
	"github.com/gin-gonic/gin"
)

type TemperatureByDateController struct {
	getTemperatureByDate *UseCase.GetTemperatureByDate
}

func NewTemperatureByDateController(getTemperatureByDate *UseCase.GetTemperatureByDate) *TemperatureByDateController{
	return &TemperatureByDateController{getTemperatureByDate: getTemperatureByDate}
}

func (controller *TemperatureByDateController) Execute(ctx *gin.Context) {
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

	// Llamada al UseCase
	temperatureUser, err := controller.getTemperatureByDate.Execute(idUs, date)
	if err != nil {
		// Manejo de error si hay un problema al obtener los datos
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respuesta exitosa con los datos obtenidos
	ctx.JSON(http.StatusOK, gin.H{"data": temperatureUser})
}
