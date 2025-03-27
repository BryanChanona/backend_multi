package UseCase

import "github.com/BryanChanona/backend_multi/src/Oxygen/domain"

type SaveOxygenUc struct {
	db domain.IOxygenRepository
}

func NewSaveOxygenUc(db domain.IOxygenRepository) *SaveOxygenUc{
	return &SaveOxygenUc{db:db}
}

func (useCase *SaveOxygenUc) Execute(oxygen domain.OxygenModel) error{
	return useCase.db.SaveOxygen(oxygen)
}