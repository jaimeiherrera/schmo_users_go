package api

import (
	"github.com/jaimeiherrera/schmo_users_go/pkg/db"
	"github.com/jaimeiherrera/schmo_users_go/src/adapter"
	"github.com/jaimeiherrera/schmo_users_go/src/gateway"
	"github.com/jaimeiherrera/schmo_users_go/src/usecase"
)

type Components struct {
	UserUserCase usecase.UserUseCaseInterface
}

func NewComponents() *Components {
	db := db.NewLocalDB()
	userRepo := getUserRepository(db)
	return &Components{
		UserUserCase: getUserUseCase(userRepo),
	}
}

func getUserUseCase(userRepo gateway.UserRepository) usecase.UserUseCaseInterface {
	return usecase.NewUserUseCase(userRepo)
}

func getUserRepository(db db.Database) gateway.UserRepository {
	return adapter.NewUserRepository(db)
}
