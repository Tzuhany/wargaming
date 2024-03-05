package config

type Config struct {
	Log     *logConf     `mapstructure:"log"`
	Service *serviceConf `mapstructure:"service"`
	JwtAuth *jwtConf     `mapstructure:"jwt"`
	Mysql   *mysqlConf   `mapstructure:"mysql"`
	Redis   *redisConf   `mapstructure:"redis"`
	Consul  *consulConf  `mapstructure:"consul"`
	Jaeger  *jaegerConf  `mapstructure:"jaeger"`
}

type serviceConf struct {
	Name string   `mapstructure:"name"`
	Addr []string `mapstructure:"addr"`
}

type mysqlConf struct {
	Addr     string `mapstructure:"addr"`
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Charset  string `mapstructure:"charset"`
}

type redisConf struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
}

type consulConf struct {
	Addr string `mapstructure:"addr"`
}

type jaegerConf struct {
	Addr string `mapstructure:"addr"`
}

type jwtConf struct {
	AccessSecret string `mapstructure:"accessSecret"`
	AccessExpire int64  `mapstructure:"accessExpire"`
}

type logConf struct {
	Level int `mapstructure:"level"`
}
