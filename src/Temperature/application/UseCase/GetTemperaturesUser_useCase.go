package UseCase

import "github.com/BryanChanona/backend_multi/src/Temperature/domain"

type GetTemperatureUser struct {
	db domain.ITemperature
}

func NewGetTemperatureUser(db domain.ITemperature) *GetTemperatureUser{
	return &GetTemperatureUser{db: db}
}

func (useCase *GetTemperatureUser)Execute()([]domain.UserTemperature, error){
	return useCase.db.GetTemperature()
}