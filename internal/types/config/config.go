package config

type Config struct {
	HTTP        *HTTP  `yaml:"http"`
	PostgresDsn string `yaml:"postgres_dsn"`
}

type HTTP struct {
	Port string `yaml:"server_port"`
}
