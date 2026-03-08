package dto

type CreateArtistRequest struct {
	Name string `json:"name;"`

	Description *string `json:"description;"`
	Genres      *string `json:"genres;"`
}

type UpdateArtistRequest = CreateArtistRequest
