package UseCase

import (
	"github.com/BryanChanona/backend_multi/src/HeartRate/domain"
	
)

type SaveHeartRateUc struct {
	db domain.IHeartRateRepository
}

func NewSaveHeartRateUc(db domain.IHeartRateRepository) *SaveHeartRateUc{
	return &SaveHeartRateUc{db: db}
}

func (useCase *SaveHeartRateUc) Execute(heartRate domain.HeartRate) error{
	return useCase.db.SaveHeartRate(heartRate)
}