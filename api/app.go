package api

import (
	"github.com/jaimeiherrera/schmo_users_go/pkg/db"
	"github.com/jaimeiherrera/schmo_users_go/src/adapter"
	"github.com/jaimeiherrera/schmo_users_go/src/gateway"
)

type Components struct {
	UserRepository gateway.UserRepository
}

func NewComponents() *Components {
	db := db.NewLocalDB()
	return &Components{
		UserRepository: getUserRepository(db),
	}
}

func getUserRepository(db db.Database) gateway.UserRepository {
	return adapter.NewUserRepository(db)
}
