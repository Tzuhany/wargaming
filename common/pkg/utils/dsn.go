package utils

import (
	"common/config"
	"github.com/cloudwego/kitex/pkg/klog"
	"strings"
)

func GetMysqlDSN() string {
	if config.Conf.Mysql == nil {
		klog.Fatal("config not found")
	}

	dsn := strings.Join([]string{config.Conf.Mysql.User, ":", config.Conf.Mysql.Password, "@tcp(", config.Conf.Mysql.Addr, ")/", config.Conf.Mysql.Database, "?charset=" + config.Conf.Mysql.Charset + "&parseTime=true"}, "")

	return dsn
}
