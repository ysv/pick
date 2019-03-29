package cmd

import (
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"github.com/ysv/pick/app/api"
	"github.com/ysv/pick/app"
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
			Value:  ":8080",
		},
	},
}

func server(c *cli.Context){
	app.GetLogger().Info("Starting pick HTTP server...")
	m := mux.NewRouter()

	// Add API subrouter with /pick & /health endpoints.
	api.RegisterRoutes(m.PathPrefix("/api").Subrouter())

	app.GetLogger().Fatal(http.ListenAndServe(":8008", m))
}
