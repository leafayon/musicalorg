package album

import (
	"time"

	"github.com/leafayon/musicalorg/internal/models/entities/artist"
	"gorm.io/gorm"
)

type Album struct {
	gorm.Model

	Artists []artist.Artist `gorm:"many2many:album_artists"`

	Name string `gorm:"unique:true"`

	Description *string
	Genres      *string

	ReleasedAt time.Time
}
