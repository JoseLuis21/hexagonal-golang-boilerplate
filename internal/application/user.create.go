package application

import (
	"context"
	"hexagonal-go/internal/domain"
)

type UserCreateService struct {
	userRepository domain.UserRepository
}

func NewUserCreateService(userRepository domain.UserRepository) UserCreateService {
	return UserCreateService{
		userRepository: userRepository,
	}
}

func (s UserCreateService) CreateUser(ctx context.Context, id, name, email, password string) error {
	user, err := domain.NewUser(id, name, email, password)
	if err != nil {
		return err
	}
	return s.userRepository.Save(ctx, user)
}
