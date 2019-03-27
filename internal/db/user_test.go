package db_test

import (
	"context"
	"git.thoughtworks.net/mahadeva/sample-golang/internal/db"
	"git.thoughtworks.net/mahadeva/sample-golang/internal/domain"
	"git.thoughtworks.net/mahadeva/sample-golang/pkg/config"
	"git.thoughtworks.net/mahadeva/sample-golang/pkg/errors"
	"git.thoughtworks.net/mahadeva/sample-golang/testutil"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var mongoClient *db.MongoDB

func TestMain(m *testing.M) {
	mongoClient = setupDB()
	os.Exit(m.Run())
}

func TestMongo_InsertUser(t *testing.T) {
	defer testutil.CleanDb()

	user := domain.NewUser("123", "Harry", "Potter")
	err := setupDB().User.Upsert(context.Background(), user)
	assert.NoError(t, err)
}

func TestMongo_FindUserByUserIDShouldReturnUser(t *testing.T) {
	defer testutil.CleanDb()

	user := domain.NewUser("123", "Harry", "Potter")
	err := setupDB().User.Upsert(context.Background(), user)

	merhchantDoc, err := mongoClient.User.FindByID(context.Background(), user.ID)

	assert.NoError(t, err)
	assert.NotNil(t, merhchantDoc)
	assert.Equal(t, user.ID, merhchantDoc.ID)
	assert.Equal(t, user.FirstName, merhchantDoc.FirstName)
}

func TestMongo_FindUserByUserIDShouldThrowErrorIfNotFound(t *testing.T) {
	defer testutil.CleanDb()

	userID := "user_id"
	user := domain.User{ID: userID, FirstName: "starbucks"}

	merhchantDoc, err := mongoClient.User.FindByID(context.Background(), user.ID)

	assert.Error(t, err)
	assert.Equal(t, errors.NewNotFoundError("User", userID), err)
	assert.Nil(t, merhchantDoc)
}

func TestMongo_UpdateUserByUserID(t *testing.T) {
	defer testutil.CleanDb()

	user := domain.NewUser("123", "Harry", "Potter")
	err := mongoClient.User.Upsert(context.Background(), user)

	user.FirstName = "New Starbucks"
	err = mongoClient.User.Upsert(context.Background(), user)
	assert.NoError(t, err)

	merhchantDoc, _ := mongoClient.User.FindByID(context.Background(), user.ID)

	assert.NotNil(t, merhchantDoc)
	assert.Equal(t, user.ID, merhchantDoc.ID)
	assert.Equal(t, user.FirstName, merhchantDoc.FirstName)
}

func setupDB() *db.MongoDB {
	return db.NewMongoClient(config.MongoURL(), config.MongoDBName())
}
