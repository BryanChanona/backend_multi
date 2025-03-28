package UseCase

import "github.com/BryanChanona/backend_multi/src/HeartRate/domain"

type UserHeartRateUc struct {
	db domain.IHeartRateRepository
}

func NewUserHeartRateUc(db domain.IHeartRateRepository)*UserHeartRateUc{
	return &UserHeartRateUc{db: db}
}

func (useCase *UserHeartRateUc) Execute() ([]domain.UserHeartRate, error){
return useCase.db.GetUserHeartRate()
}