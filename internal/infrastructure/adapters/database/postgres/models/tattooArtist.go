package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type TattooArtist struct {
	Bio           string          `gorm:"type:text;default:null"`
	Website       string          `gorm:"type:varchar(100);default:null"`
	SocialMedia   json.RawMessage `gorm:"type:jsonb;default:null"`
	UID           uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique;not null"`
	Likes         int64           `gorm:"type:int8;default:null"`
	IDCompany     uuid.UUID       `gorm:"type:uuid;not null;default:null"`
	Location      json.RawMessage `gorm:"type:json;default:null"`
	Username      string          `gorm:"type:varchar(150);not null"`
	ArtisticGenre string          `gorm:"type:varchar(150);default:null"`
	Experience    string          `gorm:"type:text;default:null"`
	Awards        json.RawMessage `gorm:"type:json;default:null"`
	Created       time.Time       `gorm:"type:timestamptz;not null"`
	Deleted       time.Time       `gorm:"type:timestamptz;default:null"`
	Updated       time.Time       `gorm:"type:timestamptz;default:null"`
	CreatedByID   uuid.UUID       `gorm:"type:uuid;default:null"`
	UpdatedByID   uuid.UUID       `gorm:"type:uuid;default:null"`
	Active        bool            `gorm:"not null"`
	DeletedByID   uuid.UUID       `gorm:"type:uuid;default:null"`
	Company       Company         `gorm:"foreignKey:IDCompany;references:UID"`
}
