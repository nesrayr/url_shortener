package postgres

type Config struct {
	Database string `config:"DB_NAME" yaml:"database"`
	User     string `config:"DB_USER" yaml:"user"`
	Password string `config:"DB_PASSWORD" yaml:"password"`
	Host     string `config:"DB_HOST" yaml:"host"`
	Port     int    `config:"DB_PORT" yaml:"port"`
	Retries  int    `config:"DB_CONNECT_RETRY" yaml:"retries"`
	PoolSize int    `config:"DB_POOL_SIZE" yaml:"pool_size"`
}
