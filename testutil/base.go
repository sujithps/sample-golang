package testutil

import (
	"git.thoughtworks.net/mahadeva/sample-golang/pkg/config"
	"git.thoughtworks.net/mahadeva/sample-golang/pkg/logger"
)

func Setup() {
	config.Load()
	logger.Setup()
}
