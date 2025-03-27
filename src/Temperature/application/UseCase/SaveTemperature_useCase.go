package UseCase

import "github.com/BryanChanona/backend_multi/src/Temperature/domain"

type SaveTemperatureUc struct {
	db domain.ITemperature
}

func NewSaveTemperatureUc(db domain.ITemperature) *SaveTemperatureUc{
 return &SaveTemperatureUc{db: db}
}

func (useCase *SaveTemperatureUc) Execute(temperature domain.Temperature) error{
	return useCase.db.SaveTemperature(temperature)
}


