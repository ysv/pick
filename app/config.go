package app

import (
	"fmt"
	"github.com/spf13/viper"
)

//type Config struct {
//	*viper.Viper
//}

func initConfig(){
	config := viper.New()

	config.SetConfigType("yml")
	config.AddConfigPath(".")
	config.SetConfigFile("config.yml")
	config.SetConfigFile("config.yaml")

	err := config.ReadInConfig()
	// TODO: Use logrus panic.
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// Check if we need this feature.
	config.WatchConfig()

	fmt.Println(config.AllKeys())
	config.Debug()

	app.config = config
}
