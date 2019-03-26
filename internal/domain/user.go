package domain

import (
	errors2 "spikes/sample-golang/pkg/errors"
)

type User struct {
	ID        string `bson:"id" json:"id"`
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
}

func NewUser(id, firstName, lastName string) *User {
	return &User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
	}
}

func (user *User) Validate() (bool, []error) {
	var validationErrors []error
	if user.ID == "" {
		validationErrors = append(validationErrors, errors2.NewRequiredFieldMisingError("ID"))
	}
	if user.FirstName == "" {
		validationErrors = append(validationErrors, errors2.NewRequiredFieldMisingError("FirstName"))
	}
	if user.LastName == "" {
		validationErrors = append(validationErrors, errors2.NewRequiredFieldMisingError("LastName"))
	}
	return len(validationErrors) == 0, validationErrors
}

type Users []User

func (users *Users) AllAreActive() bool {
	return false
}
