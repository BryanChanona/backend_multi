package domain


type ITemperature interface {
	SaveTemperature(temperature Temperature) error
	GetTemperature() ([]UserTemperature, error)
	GetTemperatureByDate(idUser int,date string,) (UserTemperature,error)
	GetTemperatureById(idUser int)([]UserTemperature, error)
}