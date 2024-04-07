package models

import (
	"time"

	"github.com/google/uuid"
)

type Seller struct {
	UID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique;not null"`
	IDUser      uuid.UUID `gorm:"type:uuid;not null"`
	Description string    `gorm:"type:text;default:null"`
	Experience  string    `gorm:"type:text;default:null"`
	DateStarted time.Time `gorm:"default:null"`
	Likes       int64     `gorm:"type:int8;default:null"`
	IDCompany   uuid.UUID `gorm:"type:uuid;not null"`
	User        User      `gorm:"foreignKey:IDUser;references:UID"`
	Company     Company   `gorm:"foreignKey:IDCompany;references:UID"`
}
