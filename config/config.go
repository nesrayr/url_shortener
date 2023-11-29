package config

import (
	"url_shortener/internal/storage/postgres"
)

type Config struct {
	Host          string `config:"HOST" yaml:"host"`
	Port          string `config:"PORT" yaml:"port"`
	TransportMode string `con`
	Database      postgres.Config
}
