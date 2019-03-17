package cmd

import (
	"os"

	"github.com/urfave/cli"
)

func Run(version string) error {
	app := cli.NewApp()

	app.Name = "Pick"
	app.Description = "lightning and reliable website analytics"
	app.Version = version

	app.Commands = []cli.Command{
		serverCommand,
		databaseCommand,
	}

	app.Run(os.Args)
	return nil
}

func beforeRun(c *cli.Context) error {
	return nil
}

func afterRun(c *cli.Context) error {
	return nil
}
