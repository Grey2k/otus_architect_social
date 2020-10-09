package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

func Load() (*Config, error) {
	cfg := new(Config)
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

type Config struct {
	Log        Log
	HTTPServer HTTPServer
	MySQL      MySQL
	JWT        JWT
}

type MySQL struct {
	User     string `env:"MYSQL_USER,required"`
	Password string `env:"MYSQL_PASSWORD,required"`
	DB       string `env:"MYSQL_DATABASE,required"`
	DBHost   string `env:"MYSQL_DB_HOST,required"`
	Port     int    `env:"MYSQL_PORT,required"`
}

func (cm *MySQL) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", cm.User, cm.Password, cm.DBHost, cm.Port, cm.DB)
}

type Log struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}

type HTTPServer struct {
	Addr         string        `env:"HTTP_SERVER_ADDR,required"`
	WriteTimeout time.Duration `env:"HTTP_SERVER_WRITETIMEOUT" envDefault:"10s"`
	ReadTimeout  time.Duration `env:"HTTP_SERVER_READTIMEOUT" envDefault:"10s"`
}

type JWT struct {
	Secret          string        `env:"JWT_SECRET,required"`
	AccessLifeTime  time.Duration `env:"JWT_ACCESS_LIFETIME,required"`
	RefreshLifeTime time.Duration `env:"JWT_REFRESH_LIFETIME,required"`
}

func (cfg *Config) Fields() log.Fields {
	return log.Fields{
		"server_addr": cfg.HTTPServer.Addr,
		"mysql_host":  cfg.MySQL.DBHost,
		"mysql_port":  cfg.MySQL.Port,
		"log_level":   cfg.Log.LogLevel,
	}
}
