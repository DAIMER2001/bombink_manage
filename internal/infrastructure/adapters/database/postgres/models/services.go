package models

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	ServiceName string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text;default:null"`
	Price       float64   `gorm:"type:numeric(10, 2);not null"`
	UID         uuid.UUID `gorm:"type:uuid;not null"`
	IDCompany   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique;not null"`
	Created     time.Time `gorm:"type:timestamptz;not null"`
	Deleted     time.Time `gorm:"type:timestamptz;default:null"`
	Updated     time.Time `gorm:"type:timestamptz;default:null"`
	CreatedByID uuid.UUID `gorm:"type:uuid;default:null"`
	UpdatedByID uuid.UUID `gorm:"type:uuid;default:null"`
	Active      bool      `gorm:"not null"`
	DeletedByID uuid.UUID `gorm:"type:uuid;default:null"`
	Company     Company   `gorm:"foreignKey:IDCompany;references:UID"`
}
