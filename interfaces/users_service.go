package interfaces

import (
	"context"
	"github.com/moisesmorillo/golang-api-example/db/dtos"
)

type UsersService interface {
	Get(ctx context.Context) (*[]dtos.Users, error)
	Create(ctx context.Context, user *dtos.Users) error
}
