package domain_test

import (
	"github.com/stretchr/testify/assert"
	"os"
	"spikes/sample-golang/internal/domain"
	errors2 "spikes/sample-golang/pkg/errors"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestUser_ValidateShouldReturnInvalidWhenNoUserID(t *testing.T) {
	user := domain.NewUser("", "Harry", "Potter")

	valid, errors := user.Validate()
	assert.False(t, valid)
	assert.Equal(t, 1, len(errors))
	assert.Equal(t, errors2.NewRequiredFieldMisingError("ID").Error(), errors[0].Error())
}

func TestUser_ValidateShouldReturnInvalidWhenNoUserName(t *testing.T) {
	user := domain.NewUser("123", "", "Potter")

	valid, errors := user.Validate()
	assert.False(t, valid)
	assert.Equal(t, 1, len(errors))
	assert.Equal(t, errors2.NewRequiredFieldMisingError("FirstName").Error(), errors[0].Error())
}

func TestUser_ValidateShouldLastName(t *testing.T) {
	user := domain.NewUser("123", "Harry", "")

	valid, errors := user.Validate()
	assert.False(t, valid)
	assert.Equal(t, 1, len(errors))
	assert.Equal(t, errors2.NewRequiredFieldMisingError("LastName").Error(), errors[0].Error())
}

func TestUser_ValidateShouldValidateAddress(t *testing.T) {
	user := domain.NewUser("123", "Harry", "Potter")

	valid, errors := user.Validate()
	assert.True(t, valid)
	assert.Equal(t, 0, len(errors))
}
