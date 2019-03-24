package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Application struct {
	config 	 *viper.Viper
	logger 	 *logrus.Logger
	database *sqlx.DB
}

var App = &app

var app Application

func InitApp(){
	initConfig()
	initLogger()
	initDatabase()
}

