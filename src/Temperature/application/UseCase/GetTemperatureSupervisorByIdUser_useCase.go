package UseCase

import "github.com/BryanChanona/backend_multi/src/Temperature/domain"

type TemperatureSupervisorUc struct {
	db domain.ITemperature
}

func NewTemperatureSupervisorUc(db domain.ITemperature) *TemperatureSupervisorUc {
	return &TemperatureSupervisorUc{db: db}
}

func (useCase *TemperatureSupervisorUc) Execute(idUser int) ([]domain.UserTemperature, error) {
	return useCase.db.GetTemperatureSupervisorByIdUser(idUser)
}