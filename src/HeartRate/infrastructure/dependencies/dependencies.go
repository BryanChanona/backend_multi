package dependencies

import (
	"log"

	"github.com/BryanChanona/backend_multi/src/HeartRate/application/UseCase"
	"github.com/BryanChanona/backend_multi/src/HeartRate/infrastructure"
	"github.com/BryanChanona/backend_multi/src/HeartRate/infrastructure/controllers"

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

func GetHeartRateController()*controllers.SaveHeartRateController{
	useCase := UseCase.NewSaveHeartRateUc(&mySQL)
	return controllers.NewSaveHeartRateController(useCase)
}
 func GetUserHeartRateController() *controllers.UserHeartRateController{
	useCase := UseCase.NewUserHeartRateUc(&mySQL)
	return controllers.NewUserHeartRateController(useCase)
 }

 func GetUserHeartRateByDateController() *controllers.UserHeartRateByDateController{
	useCase := UseCase.NewHeartRateByDateUc(&mySQL)
	return controllers.NewUserHeartRateByDateController(useCase)
 }
 func GetUserHeartRateByIdController() *controllers.UserHeartRateByIdController{
	useCase := UseCase.NewHeartRateByIdUc(&mySQL)
	return controllers.NewUserHeartRateByIdController(useCase)
 }