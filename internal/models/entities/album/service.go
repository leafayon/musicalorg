package album

import (
	"context"

	"gorm.io/gorm"
)

type Service struct {
	database *gorm.DB
}

func NewService(database *gorm.DB) *Service {
	return &Service{
		database: database,
	}
}

func (service *Service) Get(artistID uint, albumID uint) (Album, error) {
	album, err := gorm.G[Album](service.database).
		Preload("Artists", func(builder gorm.PreloadBuilder) error {
			builder.Where("id = ?", artistID)
			return nil
		}).
		Where("id = ?", albumID).
		First(context.Background())

	return album, err
}

func (service *Service) GetAllFromArtist(artistID uint) ([]Album, error) {
	album, err := gorm.G[Album](service.database).
		Preload("Artists", func(builder gorm.PreloadBuilder) error {
			builder.Where("id = ?", artistID)
			return nil
		}).
		Find(context.Background())

	return album, err
}
