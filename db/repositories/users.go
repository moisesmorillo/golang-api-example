package repositories

import (
	"context"
	"fmt"
	"github.com/moisesmorillo/golang-api-example/db/dtos"
	"github.com/moisesmorillo/golang-api-example/db/models"
	"github.com/moisesmorillo/golang-api-example/interfaces"

	"github.com/go-pg/pg/v10"
)

type userRepository struct {
	db *pg.DB
}

func NewUsersRepository(db *pg.DB) interfaces.UserRepository {
	return &userRepository{
		db,
	}
}

func (u userRepository) Get(ctx context.Context) (*[]models.Users, error) {
	users := &[]models.Users{}
	if err := u.db.
		Model(&models.Users{}).
		Context(ctx).
		Limit(10).
		Select(users); err != nil {
		return nil, fmt.Errorf("error getting users %s", err.Error())
	}

	return users, nil
}

func (u userRepository) Create(ctx context.Context, user *dtos.Users) error {
	if _, err := u.db.
		Model(user).
		Context(ctx).
		Insert(); err != nil {
		return fmt.Errorf("error creating new user %s", err.Error())
	}

	return nil
}
