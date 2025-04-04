package UseCase

import "github.com/BryanChanona/backend_multi/src/HeartRate/domain"

type HeartRateSupervisorUc struct {
	db domain.IHeartRateRepository
}

func NewHeartRateSupervisorUc(db domain.IHeartRateRepository) *HeartRateSupervisorUc {
	return &HeartRateSupervisorUc{db: db}
}

func (useCase *HeartRateSupervisorUc)Execute(idUser int)([]domain.UserHeartRate, error){
	return useCase.db.GetHeartRateSupervisorByIdUser(idUser)

}