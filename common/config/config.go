package config

type service struct {
	Name string   `mapstructure:"name" json:"name"`
	Addr []string `mapstructure:"addr" json:"addr-list"`
	LB   bool     `mapstructure:"lb" json:"lb"`
}

type redis struct {
	Addr     string `mapstructure:"addr" json:"addr"`
	Password string `mapstructure:"password" json:"password"`
	DB       int    `mapstructure:"db" json:"db"`
}

type consul struct {
	Addr string `mapstructure:"addr" json:"addr"`
}

type jaeger struct {
	Addr string `mapstructure:"host" json:"host"`
}

type jwtAuth struct {
	AccessSecret string `mapstructure:"accessSecret" json:"accessSecret"`
	AccessExpire int64  `mapstructure:"accessExpire" json:"accessExpire"`
}

type mongo struct {
	Url         string `mapstructure:"url"`
	DB          string `mapstructure:"db"`
	UserName    string `mapstructure:"userName"`
	Password    string `mapstructure:"password"`
	MinPoolSize int    `mapstructure:"minPoolSize"`
	MaxPoolSize int    `mapstructure:"maxPoolSize"`
}

type config struct {
	JwtAuth  jwtAuth             `mapstructure:"jwt" json:"jwt"`
	Redis    redis               `mapstructure:"redis" json:"redis"`
	Consul   consul              `mapstructure:"consul" json:"consul"`
	Jaeger   jaeger              `mapstructure:"jaeger" json:"jaeger"`
	Mongo    mongo               `mapstructure:"dal" json:"dal"`
	Services map[string]*service `mapstructure:"services" json:"services"`
}
