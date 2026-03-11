package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/leafayon/musicalorg/internal/models"
)

func LoadEnv(path string) error {
	err := godotenv.Load(path)

	return err
}

func Parse() (models.Config, error) {
	config, err := env.ParseAsWithOptions[models.Config](env.Options{
		Prefix: "MUSICALORG_",
	})

	return config, err
}
