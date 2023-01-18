package config

import (
	"github.com/spf13/viper"
)

type TkConfig struct {
	*Mysql  `mapstructure:"mysql"`
	*Jwt    `mapstructure:"jwt"`
	*Server `mapstructure:"server"`
	*Log    `mapstructure:"log"`
}

// NewConfig read config file from local dictionary
func NewConfig() (*TkConfig, error) {
	viper.SetConfigFile("../../setting.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	cfg := &TkConfig{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
