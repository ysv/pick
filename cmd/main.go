package cmd

import (
	"github.com/ysv/pick/app"
	"os"

	"github.com/urfave/cli"
)

func Run(version string) error {
	cliApp := cli.NewApp()

	cliApp.Name = "Pick"
	cliApp.Description = "lightning and reliable website analytics"
	cliApp.Version = version

	cliApp.Before = beforeRun
	cliApp.After = afterRun

	cliApp.Commands = []cli.Command{
		serverCommand,
		databaseCommand,
	}

	cliApp.Run(os.Args)
	return nil
}

func beforeRun(c *cli.Context) error {
	app.InitApp()
	return nil
}

func afterRun(c *cli.Context) error {
	return nil
}
