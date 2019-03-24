package app

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type RawConnectionConf map[string]string

func initDatabase(){
	conn, err := sqlx.Connect(connectionConf())
	if err != nil {
		app.logger.Fatalf("Error connecting to database: %s", err)
	}

	app.logger.Info("Successfully connected to DB")
	app.database = conn
}

func connectionConf() (string, string){
	var stringConf string

	driver := app.config.GetString("database.driver")
	rawConnectionConf := app.config.GetStringMapString("database")

	switch driver {
	case "mysql":
		stringConf = buildMySQLConnectionConf(rawConnectionConf)
	case "sqlite3":
		stringConf = buildSQLite3ConnectionConf(rawConnectionConf)
	default:
		app.logger.Fatalf("Unsupported database driver %s", driver)
	}
	return driver, stringConf
}

func buildMySQLConnectionConf(conf RawConnectionConf) string {
	mysqlConf := mysql.NewConfig()
	mysqlConf.Addr = fmt.Sprintf("%s:%s", conf["host"], conf["port"])
	mysqlConf.User = conf["user"]
	mysqlConf.Passwd = conf["password"]
	mysqlConf.Net = "tcp"
	mysqlConf.DBName = conf["name"]
	// TODO: Setup connection pool.
	return mysqlConf.FormatDSN()
}

func buildSQLite3ConnectionConf(conf RawConnectionConf) string {
	return fmt.Sprintf("%s.db", conf["name"])
}
