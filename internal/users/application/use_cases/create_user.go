package application

import "simple-go-api/internal/users/domain"

type CreateUserUseCase struct {
	repo domain.UserRepository
}

func NewCreateUserUseCase(repo domain.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo}
}

func (uc *CreateUserUseCase) Execute(id, name, email string) (*domain.User, error) {
	newUser := domain.NewUser(id, name, email)
	user, err := uc.repo.Create(newUser)

	return user, err
}
