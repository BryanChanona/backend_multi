package UseCase

import (
	"github.com/BryanChanona/backend_multi/src/Oxygen/domain"
)

type GetUserOxygenUC struct {
	db domain.IOxygenRepository
}

func NewUserOxygenUC(db domain.IOxygenRepository) *GetUserOxygenUC{
	return &GetUserOxygenUC{db: db}
}

func (useCase *GetUserOxygenUC) Execute() ([]domain.UserOxygen,error){
	return useCase.db.GetUserOxygen()
}