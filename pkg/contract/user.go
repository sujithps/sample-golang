package contract

import (
	"fmt"
	"git.thoughtworks.net/mahadeva/sample-golang/internal/domain"
)

type User struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
}

func NewUser(user *domain.User) *User {
	return &User{
		ID:          "123",
		DisplayName: fmt.Sprintf("%s, %s", user.LastName, user.FirstName),
	}
}
