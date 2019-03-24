package app

import (
	"fmt"
	"github.com/spf13/viper"
)

func initConfig(){
	config := viper.New()

	config.SetConfigType("yml")
	config.AddConfigPath(".")

	// TODO: Use default config unless file exists.
	config.SetConfigFile("config.yml")

	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	app.config = config
}
