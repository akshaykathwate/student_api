package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// Struct for HTTP server settings
type Httpserver struct {
	Addr string `yaml:"addr" env:"HTTP_ADDR" env-default:":8080"`
}

// Main Config struct embedding Httpserver
type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string     `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	Httpserver  Httpserver `yaml:"http_server"`
}

// MustLoad loads the configuration file or exits on failure
func MustLoad() *Config {
	var configPath string

	// Try to get from environment variable
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		// Fallback to command line flag
		flagPath := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		configPath = *flagPath

		if configPath == "" {
			log.Fatalf("config file path is not set via ENV or flag")
		}
	}

	// Check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	// Read and parse the config
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("cannot read config file: %s", err.Error())
	}

	return &cfg
}
