package UseCase

import "github.com/BryanChanona/backend_multi/src/Temperature/domain"

type GetTemperatureByDate struct {
	db domain.ITemperature
}

func NewGetTemperatureByDate(db domain.ITemperature) *GetTemperatureByDate{
	return &GetTemperatureByDate{db: db}
}


func (useCase *GetTemperatureByDate)Execute(id_user int,date string, ) (domain.UserTemperature,error){
	return useCase.db.GetTemperatureByDate(id_user,date)
}