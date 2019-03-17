package cmd

import (
	"fmt"
	"github.com/urfave/cli"
)

var serverCommand = cli.Command{
	Name:    "server",
	Aliases: []string{"s"},
	Usage:   "Start pick HTTP server.",
	Action:  server,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "port",
			Usage:  "server port",
			Value:  ":8080",
		},
	},
}

func server(c *cli.Context){
	fmt.Println("Starting pick HTTP server...")
}
