package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var (
	Service  *service
	Jaeger   *jaeger
	Consul   *consul
	Redis    *redis
	JwtAuth  *jwtAuth
	Mongo    *mongo
	Services map[string]*service

	runtimeViper = viper.New()
)

func Init(service string) {
	runtimeViper.SetConfigType("yaml")

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	consulAddr := os.Getenv("CONSUL_ADDR")
	if consulAddr == "" {
		panic(errors.New("not found etcd addr in env"))
	}

	Consul = &consul{Addr: consulAddr}

	err := runtimeViper.AddRemoteProvider("consul", Consul.Addr, "/config/config.yaml")
	if err != nil {
		panic(err)
	}

	err = runtimeViper.ReadRemoteConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			klog.Fatal("could not find config files")
		}
		klog.Fatal(err)
	}

	err = configMapping(service)
	if err != nil {
		klog.Fatal(err)
	}

	runtimeViper.WatchConfig()

	runtimeViper.OnConfigChange(func(e fsnotify.Event) {
		klog.Infof("config file changed: %v\n", e.String())
	})
}

func configMapping(serviceName string) error {
	c := new(config)
	if err := runtimeViper.Unmarshal(&c); err != nil {
		return fmt.Errorf("error unmarshalling config: %v", err)
	}
	Jaeger = &c.Jaeger
	Redis = &c.Redis
	JwtAuth = &c.JwtAuth
	Mongo = &c.Mongo
	Services = c.Services
	Service = c.Services[serviceName]
	return nil
}
