package db

import (
	"context"
	"git.thoughtworks.net/mahadeva/sample-golang/internal/domain"
	errors2or "git.thoughtworks.net/mahadeva/sample-golang/pkg/errors"
	"git.thoughtworks.net/mahadeva/sample-golang/pkg/profiling"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	collection = "users"
)

type UserDbClient interface {
	FindByID(ctx context.Context, userID string) (*domain.User, error)
	Upsert(ctx context.Context, userDoc *domain.User) error
}

type User struct {
	collection *mgo.Collection
}

func NewUser(db *mgo.Database) *User {
	return &User{
		collection: db.C(collection),
	}
}

func (m *User) FindByID(ctx context.Context, userID string) (*domain.User, error) {
	var userDoc domain.User
	query := bson.M{"id": userID}

	defer profiling.MongoTracer.Start(ctx, collection, "GetUserById").End()
	err := m.collection.Find(query).One(&userDoc)

	if isNotFoundErr(err) {
		return nil, errors2or.NewNotFoundError("User", userID)
	} else if err != nil {
		return nil, err
	}
	return &userDoc, err
}

func (m *User) Upsert(ctx context.Context, userDoc *domain.User) error {
	query := bson.M{"id": userDoc.ID}

	defer profiling.MongoTracer.Start(ctx, collection, "Upsert").End()
	_, err := m.collection.Upsert(query, userDoc)

	return err
}
