package UseCase

import "github.com/BryanChanona/backend_multi/src/Temperature/domain"

type TemperatureByIdUc struct {
	db domain.ITemperature
}

func NewTemperatureByIdUc(db domain.ITemperature)*TemperatureByIdUc{
	return &TemperatureByIdUc{db:db}
}

func (useCase *TemperatureByIdUc)Execute(idUser int)([]domain.UserTemperature,error){
	return useCase.db.GetTemperatureById(idUser)
}