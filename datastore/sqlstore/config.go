package sqlstore

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
)

type Config struct {
	Driver   string `default:"mysql"`
	Host     string `default:"127.0.0.1"`
	Port	 int16  `default:"3306"`
	User     string `default:"root"`
	Password string `default:""`
	Name     string `default:"pick_development"`
}

func (conf *Config) ForConnect() string {
	switch conf.Driver {
	case "mysql":
		return conf.forMySQL()
	case "sqlite3":
		return conf.forSQLite3()
	default:
		panic(fmt.Sprintf("Not Supported SQLStore driver %s", conf.Driver))
	}
}

func (conf *Config) forMySQL() string{
	mysqlConf := mysql.NewConfig()
	mysqlConf.Addr = fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	mysqlConf.User = conf.User
	mysqlConf.Passwd = conf.Password
	mysqlConf.Net = "tcp"
	mysqlConf.DBName = conf.Name
	// TODO: Setup connection pool.
	return mysqlConf.FormatDSN()
}

func (conf *Config) forSQLite3() string{
	return fmt.Sprintf("%s.db", conf.Name)
}
