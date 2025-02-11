package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	ENV               string       `yaml:"env" env-required:"true"`
	Http              HttpServer   `yaml:"http_server" env-required:"true"`
	GRPC              GRPCServer   `yaml:"grpc_server" env-required:"true"`
	ShortLinksAddress string       `yaml:"short_links_address" env-required:"true"`
	URLGenerator      URLGenerator `yaml:"url_generator" env-required:"true"`
	DB                DB
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

type URLGenerator struct {
	URLLength      int     `yaml:"url_length" env-required:"true"`
	AllowedSymbols RuneArr `yaml:"allowed_symbols" env-required:"true"`
}

// RuneArr обертка для рун, чтобы доставать их из строки cleanenv'ом
type RuneArr []rune

// UnmarshalYAML Кастомный анмаршал, чтобы доставать руны из строки
func (r *RuneArr) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	*r = []rune(str)
	return nil
}
