package domain

type IOxygenRepository interface {
	SaveOxygen(oxygen OxygenModel) error
	GetUserOxygen() ([]UserOxygen, error)
	GetOxygenByDate(date string,idUser int) ([]UserOxygen,error)
	GetOxygenById(idUser int) ([]UserOxygen, error)
}