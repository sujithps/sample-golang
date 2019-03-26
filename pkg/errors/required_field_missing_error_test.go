package errors_test

import (
	"fmt"
	errors2 "git.thoughtworks.net/mahadeva/sample-golang/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNotEmptyError(t *testing.T) {
	entity := "user name"
	notEmptyError := errors2.NewRequiredFieldMisingError(entity)

	assert.Equal(t, fmt.Sprintf("Field %s cannot be empty.", entity), notEmptyError.Error())
}
