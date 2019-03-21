package app

import (
	log "github.com/sirupsen/logrus"
	"os"
	"regexp"
)


func initLogger(){
	loggerConf := app.config.Sub("logger")
	logger := log.New()

	// TODO: File log support.
	if loggerConf.GetString("output") == "stdout"{
		logger.SetOutput(os.Stdout)
	}

	if matched, _ := regexp.MatchString(`(?i)json`, loggerConf.GetString("formatter")); matched{
		logger.SetFormatter(&log.JSONFormatter{})
	}

	switch loggerConf.GetString("level"){
	case "trace":
		logger.SetLevel(log.TraceLevel)
	case "debug":
		logger.SetLevel(log.DebugLevel)
	case "info":
		logger.SetLevel(log.InfoLevel)
	case "warn":
		logger.SetLevel(log.WarnLevel)
	case "error":
		logger.SetLevel(log.ErrorLevel)
	default:
		logger.SetLevel(log.InfoLevel)
		logger.Warn("Log level was not specified set to default: INFO")
	}

	app.logger = logger
}
