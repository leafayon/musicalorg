package router

import (
	"gorm.io/gorm"

	"github.com/leafayon/musicalorg/internal/http/controllers/api"
	"github.com/leafayon/musicalorg/internal/models/entities/artist"
)

func (router *Router) SetupRoutes(database *gorm.DB) {
	// Services
	artistService := artist.NewService(database)

	// Controllers
	artistController := api.NewArtistController(artistService)

	// Artist Routes
	router.mux.HandleFunc("GET /artists", artistController.GetAllArtists)
	router.mux.HandleFunc("GET /artists/{id}", artistController.GetArtist)
}
