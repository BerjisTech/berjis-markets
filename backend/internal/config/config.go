package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

const (
	defaultHTTPPort          = 8080
	defaultShutdownTimeout   = 15 * time.Second
	defaultReadHeaderTimeout = 5 * time.Second
	defaultEnv               = "local"
)

// Config represents runtime configuration for the API process.
type Config struct {
	AppName           string        `mapstructure:"APP_NAME"`
	Env               string        `mapstructure:"ENV"`
	HTTPPort          int           `mapstructure:"PORT"`
	ShutdownTimeout   time.Duration `mapstructure:"SHUTDOWN_TIMEOUT"`
	ReadHeaderTimeout time.Duration `mapstructure:"READ_HEADER_TIMEOUT"`
	LogLevel          string        `mapstructure:"LOG_LEVEL"`
}

// Load reads configuration from environment variables and optional config files.
func Load() (Config, error) {
	v := viper.New()
	v.SetEnvPrefix("MARKETS")
	v.AutomaticEnv()

	v.SetDefault("APP_NAME", "markets-api")
	v.SetDefault("ENV", defaultEnv)
	v.SetDefault("PORT", defaultHTTPPort)
	v.SetDefault("SHUTDOWN_TIMEOUT", defaultShutdownTimeout)
	v.SetDefault("READ_HEADER_TIMEOUT", defaultReadHeaderTimeout)
	v.SetDefault("LOG_LEVEL", "info")

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return Config{}, fmt.Errorf("config: failed to unmarshal: %w", err)
	}

	return cfg, nil
}
