package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"

	"github.com/leafayon/musicalorg/internal/models"
)

func Load(envPath string) (models.Config, error) {
	_ = godotenv.Load(envPath)

	config, err := env.ParseAsWithOptions[models.Config](env.Options{
		Prefix: "MUSICALORG_",
	})

	return config, err
}
