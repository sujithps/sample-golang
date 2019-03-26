package userservice

import (
	"context"
	"github.com/pkg/errors"
	"spikes/sample-golang/internal/db"
	"spikes/sample-golang/internal/domain"
	errors2 "spikes/sample-golang/pkg/errors"
)

type Client interface {
	Upsert(ctx context.Context, user *domain.User) error
	GetByUserID(ctx context.Context, userID string) (*domain.User, error)
}

type UserService struct {
	userStore db.UserDbClient
}

func NewUserService(userStore db.UserDbClient) *UserService {
	return &UserService{
		userStore: userStore,
	}
}

func (service *UserService) Upsert(ctx context.Context, user *domain.User) error {
	valid, validationErrors := user.Validate()
	if !valid {
		return errors2.NewValidationError(validationErrors)
	}
	err := service.userStore.Upsert(ctx, user)
	if err != nil {
		return errors.Wrap(err, "[Mongo Error] while upserting user")
	}
	return nil
}

func (service *UserService) GetByUserID(ctx context.Context, userID string) (*domain.User, error) {
	user, err := service.userStore.FindByID(ctx, userID)

	if err != nil {
		return nil, err
	}
	return user, nil
}
