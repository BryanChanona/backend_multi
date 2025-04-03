package UseCase

import "github.com/BryanChanona/backend_multi/src/Oxygen/domain"

type GetOxygenByDateUc struct {
	db domain.IOxygenRepository
}

func NewGetOxygenByDate(db domain.IOxygenRepository) *GetOxygenByDateUc {
	return &GetOxygenByDateUc{db: db}
}

func (useCase *GetOxygenByDateUc) Execute(date string,idUser int) ([]domain.UserOxygen, error) {
	return useCase.db.GetOxygenByDate(date,idUser)
}