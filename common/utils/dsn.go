package utils

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"strings"
	"wargaming/config"
)

func GetMysqlDSN() string {
	if config.Mysql == nil {
		klog.Fatal("config not found")
	}

	dsn := strings.Join([]string{config.Mysql.User, ":", config.Mysql.Password, "@tcp(", config.Mysql.Addr, ")/", config.Mysql.Database, "?charset=" + config.Mysql.Charset + "&parseTime=true"}, "")

	return dsn
}

func GetMQUrl() string {
	if config.RabbitMQ == nil {
		klog.Fatal("config not found")
	}

	url := strings.Join([]string{"amqp://", config.RabbitMQ.User, ":", config.RabbitMQ.Password, "@", config.RabbitMQ.Addr, "/"}, "")

	return url
}
