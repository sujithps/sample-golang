package dependency

import (
	"git.thoughtworks.net/mahadeva/sample-golang/internal/db"
	"git.thoughtworks.net/mahadeva/sample-golang/internal/dependency/providers"
	"git.thoughtworks.net/mahadeva/sample-golang/internal/service/userservice"
	"git.thoughtworks.net/mahadeva/sample-golang/pkg/config"
	"github.com/newrelic/go-agent"
)

type Container struct {
	newRelicApp newrelic.Application
	userService userservice.Client
	mongoDB     *db.MongoDB
}

func (container *Container) GetNewRelic() newrelic.Application {
	return container.newRelicApp
}

func (container *Container) GetUserService() userservice.Client {
	return container.userService
}

func (container *Container) Destroy() {
	container.mongoDB.Close()
}

func (container *Container) GetDBConnection() *db.MongoDB {
	return container.mongoDB
}

func Init(provider providers.NewRelicAppProvider) *Container {
	mongoClient := db.NewMongoClient(config.MongoURL(), config.MongoDBName())

	userService := userservice.NewUserService(mongoClient.User)

	return &Container{
		newRelicApp: provider(),
		userService: userService,
		mongoDB:     mongoClient,
	}
}
