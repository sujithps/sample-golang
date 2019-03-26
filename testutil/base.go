package testutil

import (
	"spikes/sample-golang/pkg/config"
	"spikes/sample-golang/pkg/logger"
)

func Setup() {
	config.Load()
	logger.Setup()
}
