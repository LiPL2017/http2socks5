package config

import (
	"github.com/spf13/viper"
)

var (
	appConfig Config
)

type Config struct {
	Listen   string `json:"listen" yaml:"listen"`
	Protocol string `json:"protocol" yaml:"protocol"`
	PemPath  string `json:"pempath" yaml:"pempath"`
	KeyPath  string `json:"keypath" yaml:"keypath"`
	Socks5   Socks5 `json:"socks5" yaml:"socks5"`
}

type Socks5 struct {
	Host     string `json:"host" yaml:"host"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
}

func loadConfig() (*Config, error) {
	if err := viper.Unmarshal(&appConfig); err != nil {
		return nil, err
	}

	return &appConfig, nil
}

func InitConfig(conf, kind string) (*Config, error) {
	viper.SetConfigType(kind)
	viper.SetConfigFile(conf)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.WatchConfig()

	return loadConfig()
}
