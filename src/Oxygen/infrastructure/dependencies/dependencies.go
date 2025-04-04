package dependencies

import (
	"log"

	"github.com/BryanChanona/backend_multi/src/Oxygen/application/UseCase"
	"github.com/BryanChanona/backend_multi/src/Oxygen/infrastructure"
	"github.com/BryanChanona/backend_multi/src/Oxygen/infrastructure/controller"
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

func GetSaveOxygenController() *controller.SaveOxygenController{
	useCase := UseCase.NewSaveOxygenUc(&mySQL)
	return controller.NewSaveOxygenController(useCase)
}
func GetOxygenationsController() *controller.UserOxygenController{
	useCase := UseCase.NewUserOxygenUC(&mySQL)
	return controller.NewUserOxygenController(useCase)
}
func GetOxygenUserByDateController()*controller.OxygenByDateController{
	useCase := UseCase.NewGetOxygenByDate(&mySQL)
	return controller.NewOxygenByDateController(useCase)
}
func GetOxygenByIdController()*controller.OxygenByIdController{
	useCase := UseCase.NewOxygenByIdUc(&mySQL)
	return controller.NewOxygenByIdController(useCase)
}

func GetOxygenSupervisorByIdUserController()*controller.OxygenSupervisorController{
	useCase := UseCase.NewOxygenSupervisorUc(&mySQL)
	return controller.NewOxygenSupervisorController(useCase)
}
