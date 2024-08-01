package gateway

import (
	entities "github.com/jaimeiherrera/schmo_users_go/src/entity"
)

type UserRepository interface {
	FindAll() ([]entities.User, error)
	FindByID(uuid string) (entities.User, error)
	Create(user entities.User) (entities.User, error)
	Update(key string, user entities.User) (entities.User, error)
	Delete(uuid string) error
}
