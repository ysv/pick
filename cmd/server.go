package cmd

import (
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"github.com/ysv/pick/app"
	"github.com/ysv/pick/app/api"
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

	// Add assets directory serving with pick.js.
	m.Handle("/assets/pick.js", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))

	app.GetLogger().Fatal(http.ListenAndServe(c.String("port"), m))
}
