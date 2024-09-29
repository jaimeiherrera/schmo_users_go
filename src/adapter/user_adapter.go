package adapter

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/jaimeiherrera/schmo_users_go/pkg/db"
	"github.com/jaimeiherrera/schmo_users_go/src/entity"
	"github.com/jaimeiherrera/schmo_users_go/src/gateway"
)

type UserAdapter struct {
	DB db.Database
}

func NewUserRepository(database db.Database) gateway.UserRepository {
	return &UserAdapter{
		DB: database,
	}
}

func (ua *UserAdapter) FindAll(ctx context.Context) ([]entity.User, error) {
	users := []entity.User{}

	data, err := ua.DB.GetAll()
	if err != nil {
		return users, err
	}

	for _, v := range data {
		userByte, err := json.Marshal(v)
		if err != nil {
			return users, err
		}

		var user entity.User
		if err := json.Unmarshal(userByte, &user); err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (ua *UserAdapter) FindByID(ctx context.Context, uuid string) (entity.User, error) {
	userEntity := entity.User{}
	user, err := ua.DB.Get(uuid)
	if err != nil {
		return entity.User{}, err
	}

	userByte, err := json.Marshal(user)
	if err != nil {
		return entity.User{}, err
	}

	if err := json.Unmarshal(userByte, &userEntity); err != nil {
		return entity.User{}, err
	}

	return userEntity, nil
}

func (ua *UserAdapter) Create(ctx context.Context, user entity.User) (entity.User, error) {
	uuid := uuid.New()

	userMap := map[string]interface{}{}
	userByte, err := json.Marshal(user)
	if err != nil {
		return entity.User{}, err
	}

	if err := json.Unmarshal(userByte, &userMap); err != nil {
		return entity.User{}, err
	}

	if err := ua.DB.Set(uuid.String(), userMap); err != nil {
		return entity.User{}, err
	}

	user.UUID = uuid.String()
	return user, nil
}

func (ua *UserAdapter) Update(ctx context.Context, key string, user entity.User) (entity.User, error) {
	userGet, err := ua.DB.Get(key)
	if err != nil {
		return entity.User{}, err
	}

	if userGet == nil {
		return entity.User{}, errors.New("user not found")
	}

	userMap := map[string]interface{}{}
	userByte, err := json.Marshal(user)
	if err != nil {
		return entity.User{}, err
	}

	if err := json.Unmarshal(userByte, &userMap); err != nil {
		return entity.User{}, err
	}

	if err := ua.DB.Set(key, userMap); err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (ua *UserAdapter) Delete(ctx context.Context, uuid string) error {
	if err := ua.DB.Delete(uuid); err != nil {
		return err
	}
	return nil
}
