package interfaces

import (
	"context"
	"github.com/moisesmorillo/golang-api-example/db/dtos"
	"github.com/moisesmorillo/golang-api-example/db/models"
)

type UserRepository interface {
	Get(ctx context.Context) (*[]models.Users, error)
	Create(ctx context.Context, user *dtos.Users) error
}
