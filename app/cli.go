package app

import (
	"github.com/sujithps/sample-golang/internal/dependency"
	"github.com/urfave/cli"
	"fmt"
)

func GetCommands(container *dependency.Container) []cli.Command {
	return []cli.Command{
		{
			Name:        "migration",
			Description: "run a migration here",
			Action: func(c *cli.Context) {
				fmt.Println(foo)
				// Run some migrations here
			},
		},
	}
}
