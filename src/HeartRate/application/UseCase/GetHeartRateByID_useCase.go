package UseCase

import "github.com/BryanChanona/backend_multi/src/HeartRate/domain"

type HeartRateByIdUc struct {
	db domain.IHeartRateRepository
}
func NewHeartRateByIdUc(db domain.IHeartRateRepository) *HeartRateByIdUc{
	return &HeartRateByIdUc{db: db}
}

func (useCase *HeartRateByIdUc) Execute(idUser int) ([]domain.UserHeartRate, error){
	return useCase.db.GetHeartRateById(idUser)
}