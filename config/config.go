package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	HTTP  `yaml:"http_server"`
	Auth  `yaml:"auth"`
	MySQL `yaml:"mysql"`
	GRPC  GRPC `yaml:"grpc"`
}

func NewConfig() *Config {
	configPath := "./config/config.yaml"

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("Ошибка при открытии файла: %s", err)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("Ошибка при прочтении конфига: %s", err)
	}

	return &cfg
}

type GRPC struct {
	Port string `yaml:"port"`
}

type HTTP struct {
	Address string `yaml:"address" env-default:"0.0.0.0:8080"`
}

type Auth struct {
	ClientID     string `yaml:"client_id" env-required:"true"`
	ClientSecret string `yaml:"client_secret" env-required:"true"`
	RedirectURI  string `yaml:"redirect_uri" env-required:"true"`
}

type MySQL struct {
	DBHost     string `yaml:"DB_HOST" env-required:"true"`
	DBDriver   string `yaml:"DB_DRIVER" env-required:"true"`
	APISecret  string `yaml:"API_SECRET" env-required:"true"`
	DBUser     string `yaml:"DB_USER" env-required:"true"`
	DBPassword string `yaml:"DB_PASSWORD" env-required:"true"`
	DBName     string `yaml:"DB_NAME" env-required:"true"`
	DBPort     int    `yaml:"DB_PORT" env-required:"true"`
}

func (cfg *Config) GetMySQLConfig() MySQL {
	return cfg.MySQL
}
