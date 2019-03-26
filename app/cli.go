package app

import (
	"github.com/urfave/cli"
	"spikes/sample-golang/internal/dependency"
)

func GetCommands(container *dependency.Container) []cli.Command {
	return []cli.Command{
		{
			Name:        "migration",
			Description: "run a migration here",
			Action: func(c *cli.Context) {
				// Run some migrations here
			},
		},
	}
}
