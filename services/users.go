package services

import (
	"context"
	"github.com/moisesmorillo/golang-api-example/db/dtos"
	"github.com/moisesmorillo/golang-api-example/db/models"
	"github.com/moisesmorillo/golang-api-example/interfaces"

	"github.com/thoas/go-funk"
)

// TODO this is a comment
type userService struct {
	repo interfaces.UserRepository
}

func NewUsersService(repo interfaces.UserRepository) interfaces.UsersService {
	return &userService{
		repo,
	}
}

func (u userService) Get(ctx context.Context) (*[]dtos.Users, error) {
	users, err := u.repo.Get(ctx)

	if err != nil {
		return nil, err
	}

	res := funk.Map(*users, func(u models.Users) dtos.Users {
		return dtos.Users{ID: u.ID, Name: u.Name}
	}).([]dtos.Users)

	return &res, nil
}

func (u userService) Create(ctx context.Context, user *dtos.Users) error {
	return u.repo.Create(ctx, user)
}
