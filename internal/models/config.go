package models

type Config struct {
	Host string `env:"HOST" envDefault:"localhost"`
	Port int    `env:"PORT" envDefault:"80"`

	Database DatabaseConfig `envPrefix:"DATABASE_"`
	Logger   LoggerConfig   `envPrefix:"LOGGER_"`
}

type DatabaseConfig struct {
	Path string `env:"PATH" envDefault:"./"`
}

type LoggerConfig struct {
	Level string `env:"LEVEL" envDefault:"INFO"`
}
