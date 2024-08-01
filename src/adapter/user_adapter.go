package adapter

import (
	"encoding/json"

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

func (ua *UserAdapter) FindAll() ([]entity.User, error) {
	users := []entity.User{}

	data, err := ua.DB.GetAll()
	if err != nil {
		return nil, err
	}

	for k, v := range data {
		userByte, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		var user entity.User
		if err := json.Unmarshal(userByte, &user); err != nil {
			return nil, err
		}

		user.UUID = k
		users = append(users, user)
	}

	return users, nil
}

func (ua *UserAdapter) FindByID(uuid string) (entity.User, error) {
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

func (ua *UserAdapter) Create(user entity.User) (entity.User, error) {
	uuid := uuid.New()
	err := ua.DB.Set(uuid.String(), user)
	if err != nil {
		return entity.User{}, err
	}

	user.UUID = uuid.String()
	return user, nil
}

func (ua *UserAdapter) Update(key string, user entity.User) (entity.User, error) {
	err := ua.DB.Set(key, user)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (ua *UserAdapter) Delete(uuid string) error {
	err := ua.DB.Delete(uuid)
	if err != nil {
		return err
	}
	return nil
}
