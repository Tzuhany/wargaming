package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var Conf *Config

func InitConfig(configFile string) {
	Conf = new(Config)

	v := viper.New()
	v.SetConfigFile(configFile)
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		log.Println("config file change")
		err := v.Unmarshal(&Conf)
		if err != nil {
			panic(err)
		}
	})

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
		return
	}

	err = v.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
}
