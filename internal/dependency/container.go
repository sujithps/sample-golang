package dependency

import (
	"github.com/newrelic/go-agent"
	"spikes/sample-golang/internal/db"
	"spikes/sample-golang/internal/dependency/providers"
	"spikes/sample-golang/internal/service/userservice"
	"spikes/sample-golang/pkg/config"
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
