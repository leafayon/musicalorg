package artist

import "gorm.io/gorm"

type Artist struct {
	gorm.Model

	Name string `gorm:"unique:true"`

	Description *string
	Genres      *string
}
