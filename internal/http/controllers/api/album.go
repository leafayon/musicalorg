package api

import (
	"github.com/leafayon/musicalorg/internal/models/entities/album"
)

type AlbumController struct {
	service *album.Service
}

func NewAlbumController(service *album.Service) *AlbumController {
	return &AlbumController{
		service: service,
	}
}
