package UseCase

import "github.com/BryanChanona/backend_multi/src/User/domain"

type RegisterUserUC struct {
	db domain.IUserRepository
}

func NewRegisterUserUC(db domain.IUserRepository) *RegisterUserUC{
	return &RegisterUserUC{db: db}
}

func (useCase *RegisterUserUC) Execute(user domain.User) error{
	return useCase.db.RegisterUser(user)
}