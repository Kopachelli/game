package conf

import (
	"github.com/jinzhu/configor"
	"github.com/rs/zerolog/log"
)

type Config struct {
	DB struct {
		User     string `required:"true" env:"DBUser"`
		Password string `required:"true" env:"DBPassword"`
		Port     int    `default:"5432"  env:"DBPort"`
		Host     string `required:"true" env:"DBHost"`
		Name     string `required:"true" env:"DBName"`
		PoolSize int    `default:"10"    env:"DBPoolSize"`
	}
	Debug        bool  `default:"false" env:"DEBUG"`
	TickInterval uint8 `default:"10" env:"TICK_INTERVAL"`
}

//New config instance
func New() *Config {
	c := &Config{}
	if err := configor.New(&configor.Config{}).Load(c); err != nil {
		log.Fatal().Err(err).Msg("Config validation error")
	}

	return c
}
