package adapter

import (
	"testing"

	"github.com/jaimeiherrera/schmo_users_go/pkg/db"
	"github.com/jaimeiherrera/schmo_users_go/src/entity"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	// TODO: Terminar estos tests
	localDB := db.NewLocalDB()
	ua := NewUserRepository(localDB)

	mockUser := entity.User{
		FirstName: "test",
		LastName:  "test",
		Email:     "test",
		Password:  "test",
	}

	t.Run("Create", func(t *testing.T) {
		user, err := ua.Create(mockUser)

		assert.NoError(t, err)
		assert.NotNil(t, user)
	})
	t.Run("Update", func(t *testing.T) {
		user, err := ua.Create(mockUser)
		assert.NoError(t, err)

		user.FirstName = "test2"
		user, err = ua.Update(user.UUID, user)
		assert.NoError(t, err)
		assert.Equal(t, "test2", user.FirstName)
	})
	t.Run("FindByID", func(t *testing.T) {
		user, err := ua.Create(mockUser)
		assert.NoError(t, err)

		userFound, err := ua.FindByID(user.UUID)
		assert.NoError(t, err)
		assert.NotNil(t, userFound)
	})
	t.Run("Delete", func(t *testing.T) {
		user, err := ua.Create(mockUser)
		assert.NoError(t, err)

		err = ua.Delete(user.UUID)
		assert.NoError(t, err)

		userFound, err := ua.FindByID(user.UUID)
		assert.Error(t, err)
		assert.Equal(t, entity.User{}, userFound)
	})
	t.Run("FindAll", func(t *testing.T) {
		_, err := ua.Create(mockUser)
		assert.NoError(t, err)

		users, err := ua.FindAll()
		assert.NoError(t, err)
		assert.NotNil(t, users)
	})
}
