package app

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Application struct {
	config *viper.Viper
	logger *logrus.Logger
}

var App = &app

var app Application

func InitApp(){
	initConfig()
	initLogger()
}

