package UseCase

import "github.com/BryanChanona/backend_multi/src/Oxygen/domain"

type OxygenSupervisorUc struct {
	db domain.IOxygenRepository
}

func NewOxygenSupervisorUc(db domain.IOxygenRepository) *OxygenSupervisorUc {
	return &OxygenSupervisorUc{db: db}
}

func (useCase *OxygenSupervisorUc) Execute(idUser int) ([]domain.UserOxygen, error) {
	return useCase.db.GetOxygenSupervisorByIdUser(idUser)
}