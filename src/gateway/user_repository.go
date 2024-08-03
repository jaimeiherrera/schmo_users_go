package gateway

import (
	entities "github.com/jaimeiherrera/schmo_users_go/src/entity"
)

// mockgen -source=src/gateway/user_repository.go -destination=src/gateway/mock/user_repository.go -package=mock
type UserRepository interface {
	FindAll() ([]entities.User, error)
	FindByID(uuid string) (entities.User, error)
	Create(user entities.User) (entities.User, error)
	Update(key string, user entities.User) (entities.User, error)
	Delete(uuid string) error
}
