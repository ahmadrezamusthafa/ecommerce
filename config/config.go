package config

import (
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

type AppConfiguration struct {
	Env  string `mapstructure:"env"`
	Port int    `mapstructure:"port"`
}

type DatabaseConfiguration struct {
	Host                 string        `mapstructure:"host"`
	Port                 int           `mapstructure:"port"`
	User                 string        `mapstructure:"user"`
	Password             string        `mapstructure:"password"`
	Name                 string        `mapstructure:"name"`
	AdditionalParameters string        `mapstructure:"additional_parameters"`
	MaxOpen              int           `mapstructure:"max_open"`
	MaxIdle              int           `mapstructure:"max_idle"`
	MaxLifetime          time.Duration `mapstructure:"max_lifetime"`
}

type CacheConfiguration struct {
	Host         string        `mapstructure:"host"`
	Port         int           `mapstructure:"port"`
	Password     string        `mapstructure:"password"`
	DialTimeout  time.Duration `mapstructure:"dial_timeout"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
	MaxIdle      int           `mapstructure:"max_idle"`
	MaxActive    int           `mapstructure:"max_active"`
	MaxLifetime  time.Duration `mapstructure:"max_lifetime"`
}

type Configuration struct {
	App      AppConfiguration      `mapstructure:"app"`
	Database DatabaseConfiguration `mapstructure:"database"`
	Cache    CacheConfiguration    `mapstructure:"cache"`
}

var configuration *Configuration
var once sync.Once

func GetConfig() *Configuration {
	once.Do(func() {
		viper.SetConfigFile("config.yaml")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}

		if err := viper.Unmarshal(&configuration); err != nil {
			log.Fatalf("Unable to decode into struct: %v", err)
		}
	})

	return configuration
}
