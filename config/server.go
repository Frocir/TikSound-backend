package config

type Server struct {
	Name string `mapstructure:"name"`
	Addr string `mapstructure:"addr"`
	Mode string `mapstructure:"mode"`
}
