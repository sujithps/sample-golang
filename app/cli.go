package app

import (
	"git.thoughtworks.net/mahadeva/sample-golang/internal/dependency"
	"github.com/urfave/cli"
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
