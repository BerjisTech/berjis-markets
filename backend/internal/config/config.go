package config

import (
	"fmt"
	"strings"
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
	DatabaseURL       string        `mapstructure:"DATABASE_URL"`
	RedisAddr         string        `mapstructure:"REDIS_ADDR"`
	RedisPassword     string        `mapstructure:"REDIS_PASSWORD"`
	JWTSecret         string        `mapstructure:"JWT_SECRET"`
	DBMaxConnections  int32         `mapstructure:"DB_MAX_CONNECTIONS"`
	CORSAllowed       []string      `mapstructure:"CORS_ALLOWED_ORIGINS"`
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
	v.SetDefault("DATABASE_URL", "postgres://market:market@localhost:5432/markets?sslmode=disable")
	v.SetDefault("REDIS_ADDR", "localhost:6379")
	v.SetDefault("REDIS_PASSWORD", "")
	v.SetDefault("JWT_SECRET", "local-dev-secret")
	v.SetDefault("DB_MAX_CONNECTIONS", 10)
	v.SetDefault("CORS_ALLOWED_ORIGINS", "*")

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return Config{}, fmt.Errorf("config: failed to unmarshal: %w", err)
	}

	cfg.normalize()

	return cfg, nil
}

func (c *Config) normalize() {
	if len(c.CORSAllowed) == 0 {
		c.CORSAllowed = []string{"*"}
		return
	}

	if len(c.CORSAllowed) == 1 && strings.Contains(c.CORSAllowed[0], ",") {
		parts := strings.Split(c.CORSAllowed[0], ",")
		var trimmed []string
		for _, p := range parts {
			if t := strings.TrimSpace(p); t != "" {
				trimmed = append(trimmed, t)
			}
		}
		if len(trimmed) > 0 {
			c.CORSAllowed = trimmed
		}
	}
}

// AllowedOrigins returns the sanitized list of allowed CORS origins.
func (c Config) AllowedOrigins() []string {
	return append([]string(nil), c.CORSAllowed...)
}
