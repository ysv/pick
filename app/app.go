package app

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/ysv/pick/datastore"
	"github.com/ysv/pick/datastore/sqlstore"
)

type Application struct {
	config 	 *viper.Viper
	logger 	 *logrus.Logger
	database datastore.Datastore
}

var App = &app

var app Application

func InitApp(){
	initConfig()
	initLogger()
}

func GetLogger() *logrus.Logger {
	return app.logger
}

func GetDB() datastore.Datastore {
	if app.database == nil {
		rawConf := app.config.Sub("database")
		conf := sqlstore.Config{
			Driver:   rawConf.GetString("driver"),
			Host:     rawConf.GetString("host"),
			Port:     int16(rawConf.GetInt("port")),
			User:     rawConf.GetString("user"),
			Password: rawConf.GetString("password"),
			Name:     rawConf.GetString("name"),
		}
		app.database = datastore.NewSQLStore(conf)
	}
	return app.database
}
