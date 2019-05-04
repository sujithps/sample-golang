package providers

import (
	"github.com/sujithps/sample-golang/pkg/config"
	"github.com/newrelic/go-agent"
	"github.com/sirupsen/logrus"
)

type NewRelicAppProvider func() newrelic.Application

func NewRelicApp() newrelic.Application {
	newRelicConfig := config.NewRelic()
	logrus.Info("New Relic Config", newRelicConfig)
	newRelicApp, err := newrelic.NewApplication(newRelicConfig)
	if err != nil {
		logrus.Error("Error Creating New Relic App", err)
		panic(err)
	}
	logrus.Info("New Relic App", newRelicApp)
	return newRelicApp
}
