package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type AllConfig struct {
	Server Server
	Log    Log
	Api    Api
}

type Server struct {
	Port  string
	Level string
}
type Log struct {
	Level    string
	FilePath string
}
type Api struct {
	ApiPath string
	ApiHost string 
	ApiPort string 
}

func ConfigInit() *AllConfig {
	config := viper.New() 
	config.AddConfigPath("./common/config")
	config.SetConfigName("application-dev")
	config.SetConfigType("yaml")

	var configs *AllConfig
	
	err := config.ReadInConfig() 
	if err != nil {
		panic(fmt.Errorf("ReadInConfig() error: %s", err))
	}
	err = config.Unmarshal(&configs)
	if err != nil {
		panic(fmt.Errorf("Unmarshal() error: %s", err))
	}
	fmt.Println("configuration: ", configs)
	return configs 
}