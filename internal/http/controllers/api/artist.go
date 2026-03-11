package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/leafayon/musicalorg/internal/models/entities/artist"
)

type ArtistController struct {
	service *artist.Service
}

func NewArtistController(service *artist.Service) *ArtistController {
	return &ArtistController{service: service}
}

func (controller *ArtistController) GetAllArtists(responseWriter http.ResponseWriter, request *http.Request) {
	artistsEntities, err := controller.service.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	if err := json.NewEncoder(responseWriter).Encode(artistsEntities); err != nil {
		fmt.Println(err)
	}
}

func (controller *ArtistController) GetArtist(responseWriter http.ResponseWriter, request *http.Request) {
	artistID, err := strconv.ParseUint(request.PathValue("id"), 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	artistEntity, _ := controller.service.Get(uint(artistID))

	if err := json.NewEncoder(responseWriter).Encode(artistEntity); err != nil {
		fmt.Println(err)
	}
}
