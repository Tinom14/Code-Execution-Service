package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"project/pkg/postgres_connect"
)

type Redis struct {
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type RabbitMQ struct {
	Host      string `yaml:"host"`
	Port      uint16 `yaml:"port"`
	QueueName string `yaml:"queue_name"`
}

type HTTPConfig struct {
	Address string `yaml:"address"`
}

type AppConfig struct {
	RabbitMQ                  `yaml:"rabbit_mq"`
	HTTPConfig                `yaml:"http"`
	postgres_connect.Postgres `yaml:"postgres"`
	Redis                     `yaml:"redis"`
	MigrationsEnabled         bool `yaml:"migrations_enabled"`
}

type AppFlags struct {
	ConfigPath string `yaml:"config_path"`
}

func ParseFlags() AppFlags {
	configPath := flag.String("config", "", "Path to config")
	flag.Parse()

	return AppFlags{
		ConfigPath: *configPath,
	}
}

func MustLoad(cfgPath string, cfg any) {
	if cfgPath == "" {
		log.Fatal("Config path is not set")
	}

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist by this path: %s", cfgPath)
	}

	if err := cleanenv.ReadConfig(cfgPath, cfg); err != nil {
		log.Fatalf("error reading config: %s", err)
	}
}
