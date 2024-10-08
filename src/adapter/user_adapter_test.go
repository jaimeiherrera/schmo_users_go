package adapter

import (
	"context"
	"errors"
	"testing"

	"github.com/jaimeiherrera/schmo_users_go/pkg/db/mock"
	"github.com/jaimeiherrera/schmo_users_go/src/entity"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUser_Create(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockDatabase(ctrl)
	ua := NewUserRepository(mockDB)
	t.Run("Create", func(t *testing.T) {
		mockDB.EXPECT().Set(gomock.Any(), gomock.Any()).Return(nil)

		us, err := ua.Create(ctx, getMockUser())

		assert.NoError(t, err)
		assert.NotNil(t, us)
	})
	t.Run("Create fail", func(t *testing.T) {
		mockDB.EXPECT().Set(gomock.Any(), gomock.Any()).Return(errors.New("error"))

		us, err := ua.Create(ctx, getMockUser())

		assert.Error(t, err)
		assert.Empty(t, us)
	})
}

func TestUser_Update(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockDatabase(ctrl)
	ua := NewUserRepository(mockDB)

	t.Run("Update success", func(t *testing.T) {
		mockDB.EXPECT().Get(gomock.Any()).Return(getMockUserMap(), nil)
		mockDB.EXPECT().Set(gomock.Any(), gomock.Any()).Return(nil)

		user, err := ua.Update(ctx, getMockUser().UUID, getMockUser())
		assert.NoError(t, err)
		assert.NotNil(t, user)
	})
	t.Run("Update fail", func(t *testing.T) {
		mockDB.EXPECT().Get(gomock.Any()).Return(getMockUserMap(), nil)
		mockDB.EXPECT().Set(gomock.Any(), gomock.Any()).Return(errors.New("error"))

		userUpdated, err := ua.Update(ctx, getMockUser().UUID, getMockUser())
		assert.Error(t, err)
		assert.Empty(t, userUpdated)
	})
	t.Run("Error getting user", func(t *testing.T) {
		mockDB.EXPECT().Get(gomock.Any()).Return(nil, errors.New("error"))

		userUpdated, err := ua.Update(ctx, getMockUser().UUID, getMockUser())
		assert.Error(t, err)
		assert.Empty(t, userUpdated)
	})
	t.Run("User not found", func(t *testing.T) {
		mockDB.EXPECT().Get(gomock.Any()).Return(nil, nil)

		userUpdated, err := ua.Update(ctx, getMockUser().UUID, getMockUser())
		assert.Error(t, err)
		assert.Empty(t, userUpdated)
	})
}

func TestUser_FindByID(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockDatabase(ctrl)
	ua := NewUserRepository(mockDB)

	t.Run("FindByID success", func(t *testing.T) {
		mockDB.EXPECT().Get(gomock.Any()).Return(getMockUserMap(), nil)

		userFound, err := ua.FindByID(ctx, "mock_id")
		assert.NoError(t, err)
		assert.NotNil(t, userFound)
	})
	t.Run("FindByID marshal error", func(t *testing.T) {
		mockDB.EXPECT().Get(gomock.Any()).Return(map[string]interface{}{"invalid": make(chan int)}, nil)

		allUsers, err := ua.FindByID(ctx, "mock_id")
		assert.Error(t, err)
		assert.Empty(t, allUsers)
	})
	t.Run("FindByID unmarshal error", func(t *testing.T) {
		mockDB.EXPECT().Get(gomock.Any()).Return(map[string]interface{}{"id": 123}, nil)

		allUsers, err := ua.FindByID(ctx, "mock_id")
		assert.Error(t, err)
		assert.Empty(t, allUsers)
	})
	t.Run("FindByID fail", func(t *testing.T) {
		mockDB.EXPECT().Get(gomock.Any()).Return(nil, errors.New("error"))

		userFound, err := ua.FindByID(ctx, "mock_id")
		assert.Error(t, err)
		assert.Empty(t, userFound)
	})
}

func TestUser_Delete(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockDatabase(ctrl)
	ua := NewUserRepository(mockDB)

	t.Run("Delete", func(t *testing.T) {
		mockDB.EXPECT().Delete(gomock.Any()).Return(nil)

		err := ua.Delete(ctx, "mock_id")
		assert.NoError(t, err)
	})
	t.Run("Delete fail", func(t *testing.T) {
		mockDB.EXPECT().Delete(gomock.Any()).Return(errors.New("error"))

		err := ua.Delete(ctx, "mock_id")
		assert.Error(t, err)
	})
}

func TestUser_FindAll(t *testing.T) {
	ctx := context.Background()
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

		allUsers, err := ua.FindAll(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, allUsers)
	})
	t.Run("FindAll DB fail", func(t *testing.T) {
		mockDB.EXPECT().GetAll().Return([]map[string]interface{}{}, errors.New("error"))

		allUsers, err := ua.FindAll(ctx)
		assert.Error(t, err)
		assert.Empty(t, allUsers)
	})
	t.Run("FindAll marshal error", func(t *testing.T) {
		mockDB.EXPECT().GetAll().Return([]map[string]interface{}{{"invalid": make(chan int)}}, nil)

		allUsers, err := ua.FindAll(ctx)
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

		allUsers, err := ua.FindAll(ctx)
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

func getMockUserMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         "mock_id",
		"first_name": "test",
		"last_name":  "test",
		"email":      "test",
		"password":   "test",
	}
}
