package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config - full configuration.
	Config struct {
		App  `yaml:"app"`
		Api  `yaml:"api"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		PG   `yaml:"postgres"`
	}

	// App - config application configuration: name and version.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// Api - config for api keys and other config.
	Api struct {
		OpenweathermapKey string `env-required:"true" env:"OW_KEY"`
	}

	// HTTP - config for http server.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log - config for logger.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// PG - config for postgres.
	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `env-required:"true"                 env:"PG_URL"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config - NewConfig - cleanenv.ReadConfig: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("config - NewConfig - cleanenv.ReadEnv: %w", err)
	}

	return cfg, nil
}
