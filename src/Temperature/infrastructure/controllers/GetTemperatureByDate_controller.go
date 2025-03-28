package controllers

import (
	"net/http"
	"strconv"

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
	// Obtener la fecha y el idUsuario desde la URL
	
	idUser, err := strconv.Atoi(ctx.Param("idUser"))
	if err != nil {
		// Manejo de error si el idUser no es válido
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	date := ctx.Param("date")

	// Llamada al UseCase
	temperatureUser, err := controller.getTemperatureByDate.Execute(idUser, date)
	if err != nil {
		// Manejo de error si hay un problema al obtener los datos
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respuesta exitosa con los datos obtenidos
	ctx.JSON(http.StatusOK, gin.H{"data": temperatureUser})
}
