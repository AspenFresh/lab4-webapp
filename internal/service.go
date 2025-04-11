package internal

import (
	"context"
)

type Storage interface {
	CreateUser(ctx context.Context, user User) error
	GetUserByEmail(ctx context.Context, email string) (User, error)
}

type UserService struct {
	storage Storage
}

func NewUserService(storage Storage) *UserService {
	return &UserService{storage: storage}
}

func (s *UserService) CreateUser(ctx context.Context, user User) (User, error) {
	err := s.storage.CreateUser(ctx, user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (User, error) {
	return s.storage.GetUserByEmail(ctx, email)
}
