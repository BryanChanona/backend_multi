package domain

type IOxygenRepository interface {
	SaveOxygen(oxygen OxygenModel) error
	GetUserOxygen() ([]UserOxygen, error)
}