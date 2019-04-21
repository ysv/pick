package cmd

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/ysv/pick/app"
)

var databaseCommand = cli.Command {
	Name: "database",
	Aliases: []string{"db"},
	Usage: 	 "Run operations with database.",
	Subcommands: []cli.Command {
		{
			Name:   "migrate",
			Usage:  "Migrate database.",
			Action: databaseMigrate,
		},
		{
			Name:   "drop",
			Usage:  "Drop database.",
			Action: databaseDrop,
		},
	},

}

func databaseMigrate(c *cli.Context){
	app.GetDB().Migrate()
}

func databaseDrop(c *cli.Context){
	fmt.Println("Dropping DB...")
}
