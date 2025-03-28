package UseCase

import "github.com/BryanChanona/backend_multi/src/Oxygen/domain"

type GetOxygenByDateUc struct {
	db domain.IOxygenRepository
}

func NewGetOxygenByDate(db domain.IOxygenRepository) *GetOxygenByDateUc {
	return &GetOxygenByDateUc{db: db}
}

func (useCase *GetOxygenByDateUc) Execute(id_user int, date string) (domain.UserOxygen, error) {
	return useCase.db.GetOxygenByDate(id_user, date)
}