package UseCase

import "github.com/BryanChanona/backend_multi/src/HeartRate/domain"

type HeartRateByDateUc struct {
	db domain.IHeartRateRepository
}

func NewHeartRateByDateUc(db domain.IHeartRateRepository)*HeartRateByDateUc{
	 return &HeartRateByDateUc{db: db}
}

func (useCase *HeartRateByDateUc) Execute(idUser int,date string) (domain.UserHeartRate,error){
	return useCase.db.GetHeartRateByDate(idUser,date)
}