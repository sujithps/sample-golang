package testutil

import (
	"github.com/sujithps/sample-golang/pkg/config"
	"github.com/sujithps/sample-golang/pkg/logger"
)

func Setup() {
	config.Load()
	logger.Setup()
}
