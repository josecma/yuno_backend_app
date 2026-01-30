package usecases

import (
	"context"
	"yuno/src/modules/users/domain/repositories"
)

type CreateUserUseCase struct {
	userRepository repositories.UserRepository
}

func (uc *CreateUserUseCase) Execute(
	ctx context.Context,
	Username, Password, FirstName, LastName, Email, PhoneNumber string,
) {
}
