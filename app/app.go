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
		conf := sqlstore.Config{
			Driver: "mysql",
			Host: "127.0.0.1",
			Port: 3306,
			User: "root",
			Password: "",
			Name: "pick_development",
		}
		app.database = datastore.NewSQLStore(conf)
	}
	return app.database
}

func CreateDB() datastore.Datastore {
	return app.database
}
