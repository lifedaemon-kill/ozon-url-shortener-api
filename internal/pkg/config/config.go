package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	ENV        string     `yaml:"env" env-required:"true"`
	Http       HttpServer `yaml:"http_server" env-required:"true"`
	GRPC       GRPCServer `yaml:"grpc_server" env-required:"true"`
	URLService URLService `yaml:"url_service" env-required:"true"`
	DB         DB
}

func Load(configPath string) *Config {
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("no such file ", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal(err)
	}

	dsn := os.Getenv("POSTGRES_DSN")

	cfg.DB.DSN = dsn

	return &cfg

}

type DB struct {
	DSN string
}
type HttpServer struct {
	Address string `yaml:"address" env-required:"true"`
}
type GRPCServer struct {
	Address string `yaml:"address" env-required:"true"`
}

type URLService struct {
	URLLength int `yaml:"url_length" env-required:"true"`
}
