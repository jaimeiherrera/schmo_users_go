package usecase

import (
	"context"

	"github.com/jaimeiherrera/schmo_users_go/src/entity"
	"github.com/jaimeiherrera/schmo_users_go/src/gateway"
)

type UserUseCaseInterface interface {
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, key string, user entity.User) (entity.User, error)
	FindUserByID(ctx context.Context, uuid string) (entity.User, error)
	FindUserAll(ctx context.Context) ([]entity.User, error)
	DeleteUser(ctx context.Context, uuid string) error
}

type UserUseCase struct {
	UserAdapter gateway.UserRepository
}

func NewUserUseCase(userAdapter gateway.UserRepository) UserUseCaseInterface {
	return &UserUseCase{
		UserAdapter: userAdapter,
	}
}

func (uc UserUseCase) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	return uc.UserAdapter.Create(ctx, user)
}

func (uc UserUseCase) UpdateUser(ctx context.Context, key string, user entity.User) (entity.User, error) {
	return uc.UserAdapter.Update(ctx, key, user)
}

func (uc UserUseCase) FindUserByID(ctx context.Context, uuid string) (entity.User, error) {
	return uc.UserAdapter.FindByID(ctx, uuid)
}

func (uc UserUseCase) FindUserAll(ctx context.Context) ([]entity.User, error) {
	return uc.UserAdapter.FindAll(ctx)
}

func (uc UserUseCase) DeleteUser(ctx context.Context, uuid string) error {
	return uc.UserAdapter.Delete(ctx, uuid)
}
