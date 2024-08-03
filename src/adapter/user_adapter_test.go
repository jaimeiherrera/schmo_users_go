package adapter

import (
	"errors"
	"testing"

	"github.com/jaimeiherrera/schmo_users_go/pkg/db/mock"
	"github.com/jaimeiherrera/schmo_users_go/src/entity"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUser_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockDatabase(ctrl)
	ua := NewUserRepository(mockDB)
	t.Run("Create", func(t *testing.T) {
		mockDB.EXPECT().Set(gomock.Any(), gomock.Any()).Return(nil)

		us, err := ua.Create(getMockUser())

		assert.NoError(t, err)
		assert.NotNil(t, us)
	})
	t.Run("Create fail", func(t *testing.T) {
		mockDB.EXPECT().Set(gomock.Any(), gomock.Any()).Return(errors.New("error"))

		us, err := ua.Create(getMockUser())

		assert.Error(t, err)
		assert.Empty(t, us)
	})
}

func TestUser_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockDatabase(ctrl)
	ua := NewUserRepository(mockDB)

	t.Run("Update", func(t *testing.T) {
		mockDB.EXPECT().Set(gomock.Any(), gomock.Any()).Return(nil)

		user, err := ua.Update(getMockUser().UUID, getMockUser())
		assert.NoError(t, err)
		assert.NotNil(t, user)
	})
	t.Run("Update fail", func(t *testing.T) {
		mockDB.EXPECT().Set(gomock.Any(), gomock.Any()).Return(errors.New("error"))

		userUpdated, err := ua.Update(getMockUser().UUID, getMockUser())
		assert.Error(t, err)
		assert.Empty(t, userUpdated)
	})
}

func TestUser_FindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockDatabase(ctrl)
	ua := NewUserRepository(mockDB)

	t.Run("FindByID success", func(t *testing.T) {
		mockDB.EXPECT().Get(gomock.Any()).Return(getMockUser(), nil)

		userFound, err := ua.FindByID("mock_id")
		assert.NoError(t, err)
		assert.NotNil(t, userFound)
	})
	t.Run("FindByID marshal error", func(t *testing.T) {
		mockDB.EXPECT().Get(gomock.Any()).Return(map[string]interface{}{"invalid": make(chan int)}, nil)

		allUsers, err := ua.FindByID("mock_id")
		assert.Error(t, err)
		assert.Empty(t, allUsers)
	})
	t.Run("FindByID unmarshal error", func(t *testing.T) {
		mockDB.EXPECT().Get(gomock.Any()).Return(map[string]interface{}{"id": 123}, nil)

		allUsers, err := ua.FindByID("mock_id")
		assert.Error(t, err)
		assert.Empty(t, allUsers)
	})
	t.Run("FindByID fail", func(t *testing.T) {
		mockDB.EXPECT().Get(gomock.Any()).Return(entity.User{}, errors.New("error"))

		userFound, err := ua.FindByID("mock_id")
		assert.Error(t, err)
		assert.Empty(t, userFound)
	})
}

func TestUser_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockDatabase(ctrl)
	ua := NewUserRepository(mockDB)

	t.Run("Delete", func(t *testing.T) {
		mockDB.EXPECT().Delete(gomock.Any()).Return(nil)

		err := ua.Delete("mock_id")
		assert.NoError(t, err)
	})
	t.Run("Delete fail", func(t *testing.T) {
		mockDB.EXPECT().Delete(gomock.Any()).Return(errors.New("error"))

		err := ua.Delete("mock_id")
		assert.Error(t, err)
	})
}

func TestUser_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockDatabase(ctrl)
	ua := NewUserRepository(mockDB)

	t.Run("FindAll success", func(t *testing.T) {
		users := []map[string]interface{}{
			{
				"id":         "mock_id",
				"first_name": "test",
				"last_name":  "test",
				"email":      "test",
				"password":   "test",
			},
		}

		mockDB.EXPECT().GetAll().Return(users, nil)

		allUsers, err := ua.FindAll()
		assert.NoError(t, err)
		assert.NotNil(t, allUsers)
	})
	t.Run("FindAll DB fail", func(t *testing.T) {
		mockDB.EXPECT().GetAll().Return([]map[string]interface{}{}, errors.New("error"))

		allUsers, err := ua.FindAll()
		assert.Error(t, err)
		assert.Empty(t, allUsers)
	})
	t.Run("FindAll marshal error", func(t *testing.T) {
		mockDB.EXPECT().GetAll().Return([]map[string]interface{}{{"invalid": make(chan int)}}, nil)

		allUsers, err := ua.FindAll()
		assert.Error(t, err)
		assert.Empty(t, allUsers)
	})
	t.Run("FindAll unmarshal error", func(t *testing.T) {
		users := []map[string]interface{}{
			{
				"id":         123,
				"first_name": "test",
				"last_name":  "test",
				"email":      "test",
				"password":   "test",
			},
		}
		mockDB.EXPECT().GetAll().Return(users, nil)

		allUsers, err := ua.FindAll()
		assert.Error(t, err)
		assert.Empty(t, allUsers)
	})
}

func getMockUser() entity.User {
	return entity.User{
		FirstName: "test",
		LastName:  "test",
		Email:     "test",
		Password:  "test",
	}
}
