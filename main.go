package main

import (
	"fmt"
	"github.com/sujithps/sample-golang/app"
	"github.com/sujithps/sample-golang/pkg/logger"
)

func handleInitError() {
	if e := recover(); e != nil {
		msg := fmt.Sprintf("Failed to load the dependency due to error : %s", e)
		logger.NonContext.Error("InitErrorHandler", msg, nil)
	}
}

func main() {
	defer handleInitError()
	app.Init()
}
