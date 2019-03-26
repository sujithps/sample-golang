package userservice_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"spikes/sample-golang/internal/db/mocks"
	"spikes/sample-golang/internal/domain"
	"spikes/sample-golang/internal/service/userservice"
	errors2 "spikes/sample-golang/pkg/errors"
	"testing"
)

var userService *userservice.UserService
var userDB *mocks.UserDbClient

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func setupTest() func() {
	userDB = &mocks.UserDbClient{}
	userService = userservice.NewUserService(userDB)
	return func() {}
}

func TestShouldUpsertADocument(t *testing.T) {
	defer setupTest()()
	user := domain.NewUser("123", "Harry", "Potter")

	userDB.On("Upsert", mock.Anything, user).Return(nil)
	err := userService.Upsert(context.Background(), user)
	assert.NoError(t, err)
}

func TestUserService_UpsertShouldReturnError(t *testing.T) {
	defer setupTest()()
	user := domain.NewUser("123", "Harry", "Potter")
	userDB.On("Upsert", mock.Anything, user).Return(errors.New(""))

	err := userService.Upsert(context.Background(), user)
	assert.Error(t, err)
	assert.Equal(t, "[Mongo Error] while upserting user: ", err.Error())
}

func TestShouldReturnValidationErrorWhileUpsertingADocument(t *testing.T) {
	defer setupTest()()

	user := domain.NewUser("123", "Harry", "Potter")
	user.ID = ""
	err := userService.Upsert(context.Background(), user)
	assert.Error(t, err)
	assert.Equal(t, errors2.NewRequiredFieldMisingError("ID").Error(), err.Error())
}

func TestShouldReturnNotFoundErrorWhenGetByUserID(t *testing.T) {
	defer setupTest()()
	userID := "non existing id"
	notFoundError := errors2.NewNotFoundError("User", userID)

	userDB.On("FindByID", mock.Anything, userID).
		Return(nil, notFoundError)
	_, err := userService.GetByUserID(context.Background(), userID)
	assert.Error(t, err)
	assert.Equal(t, notFoundError, err)
}

func TestUserService_GetByUserID(t *testing.T) {
	defer setupTest()()

	user := domain.NewUser("123", "Harry", "Potter")
	userDB.On("FindByID", mock.Anything, user.ID).
		Return(user, nil)

	storedUsersInDb, err := userService.GetByUserID(context.Background(), user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, storedUsersInDb)
}
