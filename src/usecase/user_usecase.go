package usecase

import (
	"github.com/jaimeiherrera/schmo_users_go/src/entity"
	"github.com/jaimeiherrera/schmo_users_go/src/gateway"
)

type UserUseCaseInterface interface {
	CreateUser(user entity.User) (entity.User, error)
	UpdateUser(key string, user entity.User) (entity.User, error)
	FindUserByID(uuid string) (entity.User, error)
	FindUserAll() ([]entity.User, error)
	DeleteUser(uuid string) error
}

type UserUseCase struct {
	UserAdapter gateway.UserRepository
}

func NewUserUseCase(userAdapter gateway.UserRepository) UserUseCaseInterface {
	return &UserUseCase{
		UserAdapter: userAdapter,
	}
}

func (uc UserUseCase) CreateUser(user entity.User) (entity.User, error) {
	return uc.UserAdapter.Create(user)
}

func (uc UserUseCase) UpdateUser(key string, user entity.User) (entity.User, error) {
	return uc.UserAdapter.Update(key, user)
}

func (uc UserUseCase) FindUserByID(uuid string) (entity.User, error) {
	return uc.UserAdapter.FindByID(uuid)
}

func (uc UserUseCase) FindUserAll() ([]entity.User, error) {
	return uc.UserAdapter.FindAll()
}

func (uc UserUseCase) DeleteUser(uuid string) error {
	return uc.UserAdapter.Delete(uuid)
}
