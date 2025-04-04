package dependencies

import (
	"log"

	"github.com/BryanChanona/backend_multi/src/Temperature/application/UseCase"
	"github.com/BryanChanona/backend_multi/src/Temperature/infrastructure"
	"github.com/BryanChanona/backend_multi/src/Temperature/infrastructure/controllers"
	"github.com/BryanChanona/backend_multi/src/helpers"
)

var (
	mySQL infrastructure.MySQL
)

func Init() {
	db, err := helpers.ConnMySQL()

	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	mySQL = *infrastructure.NewMySQL(db)

}


func GetSaveTemperatureController() *controllers.SaveTemperatureController{
	useCase := UseCase.NewSaveTemperatureUc(&mySQL)
	return controllers.NewSaveTemperatureController(useCase)
}
func GetUserTemperaturesController() *controllers.GetTemperatureUsersController{
	useCase := UseCase.NewGetTemperatureUser(&mySQL)
	return controllers.NewGetTemperatureUsersController(useCase)
}

func GetUserTemperatureByDateController() *controllers.TemperatureByDateController{
	useCase := UseCase.NewGetTemperatureByDate(&mySQL)
	return controllers.NewTemperatureByDateController(useCase)
}
func GetUserTemperaturesByIdController()*controllers.TemperatureByIdController{
	useCase := UseCase.NewTemperatureByIdUc(&mySQL)
	return controllers.NewTemperatureByIdController(useCase)
}
func GetTemperatureSupervisorByIdUserController()*controllers.GetTemperatureSupervisorByIdUserController{
	useCase := UseCase.NewTemperatureSupervisorUc(&mySQL)
	return controllers.NewGetTemperatureSupervisorByIdUserController(useCase)
}


