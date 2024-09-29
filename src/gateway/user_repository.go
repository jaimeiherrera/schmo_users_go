package gateway

import (
	"context"

	entities "github.com/jaimeiherrera/schmo_users_go/src/entity"
)

// mockgen -source=src/gateway/user_repository.go -destination=src/gateway/mock/user_repository.go -package=mock
type UserRepository interface {
	FindAll(ctx context.Context) ([]entities.User, error)
	FindByID(ctx context.Context, uuid string) (entities.User, error)
	Create(ctx context.Context, user entities.User) (entities.User, error)
	Update(ctx context.Context, key string, user entities.User) (entities.User, error)
	Delete(ctx context.Context, uuid string) error
}
