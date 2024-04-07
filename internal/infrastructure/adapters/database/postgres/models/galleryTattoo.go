package models

import "github.com/google/uuid"

type GalleryTattoo struct {
	MediaURL     string       `gorm:"type:varchar(200);default:null"`
	UID          uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique;not null"`
	IDArtist     uuid.UUID    `gorm:"type:uuid;not null"`
	TypeMedia    string       `gorm:"type:varchar(100);default:null"`
	TattooArtist TattooArtist `gorm:"foreignKey:IDArtist;references:UID"`
}
