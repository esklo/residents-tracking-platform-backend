package user

import (
	"context"
	"github.com/esklo/residents-tracking-platform-backend/internal/model"
	"github.com/esklo/residents-tracking-platform-backend/mocks"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"testing"
)

func setup(t *testing.T) (*mocks.MockUserRepository, *mocks.MockDepartmentService, *Service, context.Context, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	// mocks
	mockUserRepository := mocks.NewMockUserRepository(ctrl)
	mockDepartmentService := mocks.NewMockDepartmentService(ctrl)
	logger := zap.NewNop()

	// service
	userSvc := NewService(mockUserRepository, mockDepartmentService, logger)

	ctx := context.Background()

	return mockUserRepository, mockDepartmentService, userSvc, ctx, ctrl
}

func TestUserCreate(t *testing.T) {
	mockUserRepository, _, userSvc, ctx, ctrl := setup(t)
	defer ctrl.Finish()

	user := &model.User{
		Email:     "test@test.ru",
		Password:  "password123",
		FirstName: "first",
		Role:      model.UserRoleAdmin,
	}

	mockUserRepository.EXPECT().Create(ctx, gomock.Any()).Return(user, nil)
	createdUser, err := userSvc.Create(ctx, user)

	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, user, createdUser)
	assert.NotEmpty(t, createdUser.Salt)
}

func TestUserGet(t *testing.T) {
	mockUserRepository, _, userSvc, ctx, ctrl := setup(t)
	defer ctrl.Finish()

	uid := uuid.New()
	user := &model.User{
		Id:        uid,
		Email:     "test@test.ru",
		Password:  "password123",
		FirstName: "first",
		Role:      model.UserRoleAdmin,
	}

	mockUserRepository.EXPECT().GetByID(ctx, gomock.Any()).Return(user, nil)
	foundUser, err := userSvc.Get(ctx, uid.String())

	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, user, foundUser)

	mockUserRepository.EXPECT().GetByID(ctx, gomock.Any()).Return(nil, nil)
	_, err = userSvc.Get(ctx, uid.String())
	assert.Error(t, err)
	assert.Equal(t, err, model.ErrorNotFound)
}

func TestUserGetAll(t *testing.T) {
	mockUserRepository, mockDepartmentService, userSvc, ctx, ctrl := setup(t)
	defer ctrl.Finish()

	uid := uuid.New()
	user := &model.User{
		Id:        uid,
		Email:     "test@test.ru",
		Password:  "password123",
		FirstName: "first",
		Role:      model.UserRoleAdmin,
	}

	users := []*model.User{user}

	mockUserRepository.EXPECT().GetAll(ctx).Return(users, nil)
	mockUserRepository.EXPECT().GetAllWithDepartmentIds(ctx, gomock.Any()).Return(nil, nil)
	mockDepartmentService.EXPECT().GetAll(ctx, gomock.Any()).Return(nil, nil)

	getAll, err := userSvc.GetAll(ctx, nil)

	assert.NoError(t, err)
	assert.NotNil(t, getAll)
	assert.Equal(t, users, getAll)

	getAllWithDepartment, err := userSvc.GetAll(ctx, &uid)

	assert.NoError(t, err)
	assert.Nil(t, getAllWithDepartment)
}

func TestUserUpdate(t *testing.T) {
	mockUserRepository, _, userSvc, ctx, ctrl := setup(t)
	defer ctrl.Finish()

	uid := uuid.New()
	user := &model.User{
		Id:        uid,
		Email:     "test@test.ru",
		Password:  "password123",
		FirstName: "first",
		Role:      model.UserRoleAdmin,
	}

	mockUserRepository.EXPECT().Update(ctx, gomock.Any()).Return(nil)
	err := userSvc.Update(ctx, user)
	assert.NoError(t, err)

	mockError := errors.New("some error")
	mockUserRepository.EXPECT().Update(ctx, gomock.Any()).Return(mockError)
	err = userSvc.Update(ctx, user)
	assert.Error(t, err)
	assert.Equal(t, mockError, err)
}
