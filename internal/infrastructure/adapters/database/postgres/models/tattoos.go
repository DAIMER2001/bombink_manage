package models

import (
	"time"

	"github.com/google/uuid"
)

type Tattoos struct {
	TattooName     string            `gorm:"type:varchar(100);not null"`
	Description    string            `gorm:"type:text;default:null"`
	UID            uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique;not null"`
	IDArtist       uuid.UUID         `gorm:"type:uuid;not null;"`
	TattooCategory string            `gorm:"type:varchar(100);default:null"`
	Created        time.Time         `gorm:"type:timestamptz;not null"`
	Deleted        time.Time         `gorm:"type:timestamptz;default:null"`
	Updated        time.Time         `gorm:"type:timestamptz;default:null"`
	CreatedByID    uuid.UUID         `gorm:"type:uuid;default:null"`
	UpdatedByID    uuid.UUID         `gorm:"type:uuid;default:null"`
	DeletedByID    uuid.UUID         `gorm:"type:uuid;default:null"`
	Likes          int64             `gorm:"type:int8;default:null"`
	Price          float64           `gorm:"type:numeric(10, 2);default:null"`
	TattooHealed   string            `gorm:"type:varchar(150);default:null"`
	TattooArtist   TattooArtist      `gorm:"foreignKey:IDArtist;references:UID"`
	Reactions      []TattooReactions `gorm:"foreignKey:IDTattoo;references:UID"`
}
