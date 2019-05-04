package errors_test

import (
	"errors"
	errors2 "github.com/sujithps/sample-golang/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNotFoundError_Error(t *testing.T) {
	notFoundError := errors2.NewNotFoundError("User", "userID")
	assert.True(t, errors2.IsNotFoundError(notFoundError))

	invalidErr := errors.New("other err")
	assert.False(t, errors2.IsNotFoundError(invalidErr))
}
