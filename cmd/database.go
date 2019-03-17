package cmd

import (
	"fmt"
	"github.com/urfave/cli"
)

var databaseCommand = cli.Command {
	Name: "database",
	Aliases: []string{"db"},
	Usage: 	 "Run operations with database.",
	Subcommands: []cli.Command {
		{
			Name:   "create",
			Usage:  "Create database.",
			Action: databaseCreate,
		},
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

func databaseCreate(c *cli.Context){
	fmt.Println("Creating DB...")
}

func databaseMigrate(c *cli.Context){
	fmt.Println("Migrating DB...")
}

func databaseDrop(c *cli.Context){
	fmt.Println("Dropping DB...")
}
