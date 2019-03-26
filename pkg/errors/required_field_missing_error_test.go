package errors_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	errors2 "spikes/sample-golang/pkg/errors"
	"testing"
)

func TestNotEmptyError(t *testing.T) {
	entity := "user name"
	notEmptyError := errors2.NewRequiredFieldMisingError(entity)

	assert.Equal(t, fmt.Sprintf("Field %s cannot be empty.", entity), notEmptyError.Error())
}
