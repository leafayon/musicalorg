package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"

	"github.com/leafayon/musicalorg/internal/models"
)

func LoadEnv(path string) error {
	return godotenv.Load(path)
}

func Parse() (models.Config, error) {
	config, err := env.ParseAsWithOptions[models.Config](env.Options{
		Prefix: "MUSICALORG_",
	})

	return config, err
}
