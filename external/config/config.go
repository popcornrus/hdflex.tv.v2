package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"go-chat/external/logger/sl"
	ucfg "go.uber.org/config"
	"go.uber.org/fx"
	"log/slog"
	"os"
	"time"
)

type (
	ResultConfig struct {
		fx.Out

		Config   *Config
		Provider ucfg.Provider
	}

	Config struct {
		Env        string `yaml:"env"`
		HTTPServer `yaml:"http_server"`
		ENVState   `yaml:"env_state"`
		MongoDB    `yaml:"mongodb"`
	}

	ENVState struct {
		Local string `yaml:"local" env-default:"local"`
		Dev   string `yaml:"dev" env-default:"dev"`
		Prod  string `yaml:"prod" env-default:"prod"`
	}

	HTTPServer struct {
		Address     string        `yaml:"address" env-default:"localhost:8080"`
		Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
		IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	}

	/*	MongoDB struct {
		URI           string `env:"MONGODB_URI" env-default:"mongodb://localhost:27017"`
		User          string `env:"MONGODB_USER" env-default:"root"`
		Password      string `env:"MONGODB_PASSWORD" env-default:"root"`
		Host          string `env:"MONGODB_HOST" env-default:"localhost"`
		Port          string `env:"MONGODB_PORT" env-default:"27017"`
		DBName        string `env:"MONGODB_DBNAME" env-default:"rust"`
		AuthDatabase  string `env:"MONGODB_AUTH_DBNAME" env-default:"admin"`
		AuthMechanism string `env:"MONGODB_AUTH_MECHANISM" env-default:"SCRAM-SHA-1"`
	}*/

	MongoDB struct {
		URI           string `yaml:"mongodb_uri" env-default:"mongodb://mongo:27017"`
		User          string `yaml:"mongodb_user" env-default:"root"`
		Password      string `yaml:"mongodb_password" env-default:"root"`
		Host          string `yaml:"mongodb_host" env-default:"localhost"`
		Port          string `yaml:"mongodb_port" env-default:"27017"`
		DBName        string `yaml:"mongodb_dbname" env-default:"rust"`
		AuthDatabase  string `yaml:"mongodb_auth_database" env-default:"admin"`
		AuthMechanism string `yaml:"mongodb_auth_mechanism" env-default:"SCRAM-SHA-1"`
	}
)

func NewConfig() (*Config, error) {
	log := slog.With(
		slog.String("op", "config.NewConfig"),
	)

	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return nil, fmt.Errorf("CONFIG_PATH is not set")
	}

	log.Info("Config path", sl.String("path", configPath))

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("cannot read config: %s", err)
	}

	return &cfg, nil
}
