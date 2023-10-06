package app

import "github.com/vnniciusg/src/internal/domain"

type UserUseCase struct {
	UserRepository domain.UserRepository
}

func NewUserUseCase(userRepository domain.UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

func (uc *UserUseCase) Create(user *domain.User) error {
	return uc.UserRepository.Create(user)
}

func (uc *UserUseCase) GetByID(id int) (*domain.User, error) {
	return uc.UserRepository.GetByID(id)
}
