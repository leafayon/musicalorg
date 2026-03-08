package artist

import (
	"context"

	"gorm.io/gorm"

	"github.com/leafayon/musicalorg/internal/models/dto"
)

type Service struct {
	database *gorm.DB
}

func NewService(database *gorm.DB) *Service {
	return &Service{
		database: database,
	}
}

func (service *Service) Get(id uint) (Artist, error) {
	artist, err := gorm.G[Artist](service.database).
		Where("id = ?", id).
		First(context.Background())

	return artist, err
}

func (service *Service) GetAll() ([]Artist, error) {
	artists, err := gorm.G[Artist](service.database).
		Find(context.Background())

	return artists, err
}

func (service *Service) Create(createArtistRequest *dto.CreateArtistRequest) error {
	err := gorm.G[Artist](service.database).
		Create(
			context.Background(),
			&Artist{
				Name:        createArtistRequest.Name,
				Description: createArtistRequest.Description,
				Genres:      createArtistRequest.Genres,
			},
		)

	return err
}

func (service *Service) Update(artistID uint, updateArtistRequest *dto.UpdateArtistRequest) (int, error) {
	rowsAffected, err := gorm.G[Artist](service.database).
		Where("id = ?", artistID).
		Updates(
			context.Background(),
			Artist{
				Name:        updateArtistRequest.Name,
				Description: updateArtistRequest.Description,
				Genres:      updateArtistRequest.Genres,
			},
		)

	return rowsAffected, err
}

func (service *Service) Delete(artistID uint) (int, error) {
	rowsAffected, err := gorm.G[Artist](service.database).
		Where("id = ?", artistID).
		Delete(context.Background())

	return rowsAffected, err
}
