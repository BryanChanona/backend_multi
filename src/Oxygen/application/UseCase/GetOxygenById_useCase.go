package UseCase

import "github.com/BryanChanona/backend_multi/src/Oxygen/domain"

type OxygenByIdUc struct {
	db domain.IOxygenRepository
}

func NewOxygenByIdUc(db domain.IOxygenRepository) *OxygenByIdUc{
	return &OxygenByIdUc{db: db}
}

func (useCase *OxygenByIdUc)Execute(idUser int)([]domain.UserOxygen,error){
	return useCase.db.GetOxygenById(idUser)
}