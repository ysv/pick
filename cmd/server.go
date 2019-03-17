package cmd

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/ysv/pick/api"
	"net/http"
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
			Value:  ":8090",
		},
	},
}

func server(c *cli.Context){
	fmt.Println("Starting pick HTTP server...")

	handler := http.NewServeMux()
	handler.Handle("/", api.Router())
	//handler.
	http.ListenAndServe(c.String("port"), handler)
}
