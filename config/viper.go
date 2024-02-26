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
	Server   *server
	Service  *service
	Mysql    *mysql
	Jaeger   *jaeger
	RabbitMQ *rabbitMq
	Consul   *consul
	Redis    *redis
	JwtAuth  *jwtAuth

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

	runtimeViper.OnConfigChange(func(e fsnotify.Event) {
		klog.Infof("config file changed: %v\n", e.String())
	})
	runtimeViper.WatchConfig()
}

func configMapping(serviceName string) error {
	c := new(config)
	if err := runtimeViper.Unmarshal(&c); err != nil {
		return fmt.Errorf("error unmarshalling config: %v", err)
	}
	Server = &c.Server
	Jaeger = &c.Jaeger
	Mysql = &c.Mysql
	Redis = &c.Redis
	JwtAuth = &c.JwtAuth
	RabbitMQ = &c.RabbitMq
	Service = GetService(serviceName)
	return nil
}

func GetService(srvName string) *service {
	addrList := runtimeViper.GetStringSlice("services." + srvName + ".addr")

	return &service{
		Name:     runtimeViper.GetString("services." + srvName + ".name"),
		AddrList: addrList,
		LB:       runtimeViper.GetBool("services." + srvName + ".load-balance"),
	}
}

func InitForTest() {
	Server = &server{
		Name: "tiktok",
	}

	Jaeger = &jaeger{
		Addr: "localhost:6831",
	}

	JwtAuth = &jwtAuth{
		AccessExpire: 3600,
		AccessSecret: "MTAxNTkwMTg1Mw==",
	}

	Consul = &consul{
		Addr: "localhost:8500",
	}

	Mysql = &mysql{
		Addr:     "localhost:3307",
		Database: "wargame",
		User:     "wargame",
		Password: "wargame",
		Charset:  "utf8mb4",
	}

	RabbitMQ = &rabbitMq{
		Addr:     "localhost:5672",
		User:     "wargame",
		Password: "wargame",
	}

	Redis = &redis{
		Addr:     "localhost:6378",
		Password: "wargame",
	}
}
