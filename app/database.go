package app

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type RawConnectionConf map[string]string

func initDatabase(){
	conn, err := sqlx.Connect("mysql", connectionConf())
	if err != nil {
		app.logger.Fatalf("Error connecting to database: %s", err)
	}
	app.database = conn
}

func connectionConf() string{
	dbConf := app.config.Sub("database")
	switch dbConf.GetString("driver") {
	case "mysql":
		return buildMySQLConnectionConf(app.config.GetStringMapString("database"))
	default:
		app.logger.Fatalf("Unsupported database driver %s", dbConf.GetString("driver"))
	}
	return ""
}

func buildMySQLConnectionConf(conf RawConnectionConf) string {
	mysqlConf := mysql.NewConfig()
	mysqlConf.Addr = fmt.Sprintf("%s:%s", conf["host"], conf["port"])
	mysqlConf.User = conf["user"]
	mysqlConf.Passwd = conf["password"]
	mysqlConf.Net = "tcp"
	mysqlConf.DBName = conf["name"]
	return mysqlConf.FormatDSN()
}
