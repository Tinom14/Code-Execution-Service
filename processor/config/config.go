package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"project/pkg/postgres_connect"
)

type RabbitMQ struct {
	Host      string `yaml:"host"`
	Port      uint16 `yaml:"port"`
	QueueName string `yaml:"queue_name"`
}

type PrometheusConfig struct {
	Port uint16 `yaml:"port"`
}

type HTTPConfig struct {
	CommitURL string `yaml:"commit_url"`
	Port      uint16 `yaml:"port"`
}

type AppConfig struct {
	RabbitMQ                  `yaml:"rabbit_mq"`
	PrometheusConfig          `yaml:"prometheus"`
	HTTPConfig                `yaml:"http"`
	postgres_connect.Postgres `yaml:"postgres"`
}

func ParseFlags() string {
	configPath := flag.String("config", "", "Path to config")
	flag.Parse()
	return *configPath
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
