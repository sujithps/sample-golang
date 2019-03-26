package errors_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	errors2 "spikes/sample-golang/pkg/errors"
	"testing"
)

func TestNotFoundError_Error(t *testing.T) {
	notFoundError := errors2.NewNotFoundError("User", "userID")
	assert.True(t, errors2.IsNotFoundError(notFoundError))

	invalidErr := errors.New("other err")
	assert.False(t, errors2.IsNotFoundError(invalidErr))
}
