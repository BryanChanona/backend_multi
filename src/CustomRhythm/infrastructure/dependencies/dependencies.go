package dependencies

import (
	"log"

	"github.com/BryanChanona/backend_multi/src/CustomRhythm/application/UseCase"
	"github.com/BryanChanona/backend_multi/src/CustomRhythm/infrastructure"
	"github.com/BryanChanona/backend_multi/src/CustomRhythm/infrastructure/controllers"
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

func GetRegisterCustomRhythm() *controllers.RegisterCustomRhythmController{
	useCase := UseCase.NewRegisterCustomRhythmUC(&mySQL)
	return controllers.NewRegisterCustomRhythmController(useCase)
}

