package domain


type IHeartRateRepository interface {
	SaveHeartRate(heartRate HeartRate) error
	GetUserHeartRate() ([]UserHeartRate, error)
	GetHeartRateByDate(idUser int,date string) (UserHeartRate,error)
	GetHeartRateById(idUser int) ([]UserHeartRate,error)
}