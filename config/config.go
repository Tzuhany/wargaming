package config

type server struct {
	Name string `mapstructure:"name" json:"name"`
}

type service struct {
	Name     string   `mapstructure:"name" json:"name"`
	AddrList []string `mapstructure:"addr-list" json:"addr-list"`
	LB       bool     `mapstructure:"lb" json:"lb"`
}

type pasetoConfig struct {
	PubKey   string `mapstructure:"pub_key" json:"pub_key"`
	Implicit string `mapstructure:"implicit" json:"implicit"`
}

type mysql struct {
	Addr     string `mapstructure:"addr" json:"addr"`
	Database string `mapstructure:"database" json:"database"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Charset  string `mapstructure:"charset" json:"charset"`
}

type redis struct {
	Addr     string `mapstructure:"addr" json:"addr"`
	Password string `mapstructure:"password" json:"password"`
}

type consul struct {
	Addr string `mapstructure:"addr" json:"addr"`
}

type rabbitMq struct {
	Addr     string `mapstructure:"addr" json:"addr"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type jaeger struct {
	Addr string `mapstructure:"host" json:"host"`
}

type jwtAuth struct {
	AccessSecret string `mapstructure:"accessSecret" json:"accessSecret"`
	AccessExpire int64  `mapstructure:"accessExpire" json:"accessExpire"`
}

type config struct {
	Server       server       `mapstructure:"server" json:"server"`
	JwtAuth      jwtAuth      `mapstructure:"jwt" json:"jwt"`
	Mysql        mysql        `mapstructure:"mysql" json:"mysql"`
	Redis        redis        `mapstructure:"redis" json:"redis"`
	RabbitMq     rabbitMq     `mapstructure:"rabbitmq" json:"rabbitmq"`
	Consul       consul       `mapstructure:"consul" json:"consul"`
	Jaeger       jaeger       `mapstructure:"jaeger" json:"jaeger"`
	PasetoConfig pasetoConfig `mapstructure:"pasetoConfig" json:"pasetoConfig"`
}
