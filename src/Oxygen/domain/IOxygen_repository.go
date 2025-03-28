package domain

type IOxygenRepository interface {
	SaveOxygen(oxygen OxygenModel) error
	GetUserOxygen() ([]UserOxygen, error)
	GetOxygenByDate(idUser int,date string,) (UserOxygen,error)
	GetOxygenById(idUser int) ([]UserOxygen, error)
}