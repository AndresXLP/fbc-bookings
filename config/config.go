package config

import (
	"log"
	"os"
	"sync"

	"github.com/andresxlp/gosuite/config"
)

type Config struct {
	Server   Server   `validate:"required" mapstructure:"server"`
	Postgres Postgres `validate:"required" mapstructure:"postgres"`
}

type Server struct {
	Host string `validate:"required" mapstructure:"host"`
	Port int    `validate:"required" mapstructure:"port"`
}

type Postgres struct {
	DbHost     string `validate:"required" mapstructure:"db_host"`
	DbPort     int    `validate:"required" mapstructure:"db_port"`
	DbUser     string `validate:"required" mapstructure:"db_user"`
	DbPassword string `validate:"required" mapstructure:"db_password"`
	DbName     string `validate:"required" mapstructure:"db_name"`
}

var (
	once sync.Once
	Cfg  Config
)

func Environments() Config {
	once.Do(func() {
		if os.Getenv("ENVIRONMENT") == "DEVELOP" {
			if err := config.SetEnvsFromFile(".env"); err != nil {
				log.Panic(err)
			}
		}

		if err := config.GetConfigFromEnv(&Cfg); err != nil {
			log.Panicf("Error parsing environment vars %v", err)
		}
	})

	return Cfg
}
