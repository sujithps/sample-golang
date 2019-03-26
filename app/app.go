package app

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/urfave/negroni"
	"os"
	"os/signal"
	"spikes/sample-golang/internal/dependency"
	"spikes/sample-golang/internal/dependency/providers"
	"spikes/sample-golang/internal/router"
	"spikes/sample-golang/pkg/config"
	"spikes/sample-golang/pkg/logger"
	"spikes/sample-golang/pkg/middleware"
	"strconv"
	"syscall"
	"time"
)

func Init() {
	config.Load()
	logger.Setup()
	container := dependency.Init(providers.NewRelicApp)
	setAppTimeToUTC()

	commands := cli.NewApp()
	commands.Name = "Sample GoLang App"
	commands.Version = "1.0.0"

	commands.Action = func(c *cli.Context) {
		startHTTPServer(container)
	}
	commands.Commands = GetCommands(container)
	if err := commands.Run(os.Args); err != nil {
		panic(err)
	}
}

func setAppTimeToUTC() {
	time.Local = time.UTC
}

func startHTTPServer(container *dependency.Container) {
	logger.NonContext.Info("startHTTPServer", "Starting Web Server", nil)

	server := negroni.New(negroni.NewRecovery())
	server.Use(middleware.JSONAPI())
	server.Use(middleware.CorrelationID())
	server.Use(middleware.Recover())
	server.Use(middleware.HTTPStatLogger())
	server.UseHandler(router.Router(container))

	port := fmt.Sprintf(":%s", strconv.Itoa(config.Port()))
	logger.NonContext.Info("startHTTPServer", "Starting Server on port "+port, nil)

	go server.Run(port)
	waitForServerShutdown(container)
}

func waitForServerShutdown(container *dependency.Container) {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		sig := <-signalChan
		switch sig {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			logger.NonContext.Info("", fmt.Sprintf("Received a signal %s", sig), nil)
			container.Destroy()
			os.Exit(0)
		default:
			logger.NonContext.Info("", fmt.Sprintf("Received a unexpected signal %s", sig), nil)
			container.Destroy()
		}
	}
}
