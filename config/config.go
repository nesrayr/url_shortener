package config

import (
	"url_shortener/internal/storage/postgres"
)

type Config struct {
	Host          string `config:"HOST" yaml:"host"`
	Port          string `config:"PORT" yaml:"port"`
	TransportMode string `config:"TRANSPORT_MODE" yaml:"transport_mode"`
	StorageMode   string `config:"STORAGE_MODE" yaml:"storage_mode"`
	Database      postgres.Config
}
