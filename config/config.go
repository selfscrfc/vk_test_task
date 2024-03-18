package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Server   Server
	Logger   Logger
	Postgres Postgres
	Redis    Redis
}

type Server struct {
	Port            string
	Version         string
	AccessSecret    string
	RefreshSecret   string
	AccessLifetime  int64
	RefreshLifetime int64
}

type Redis struct {
	Host     string
	Port     string
	Database int
}

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

type Logger struct {
	InFile bool
}

func ParseConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	var cp string
	if os.Getenv("IN_DOCKER") == "TRUE" {
		cp = "/build/config"
	} else {
		cp = "./config"
	}
	viper.AddConfigPath(cp)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
