package models

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	Name        string    `gorm:"type:varchar(70);not null"`
	UID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique;not null"`
	Logo        string    `gorm:"type:varchar(200);default:null"`
	Created     time.Time `gorm:"type:timestamptz;not null"`
	CreatedByID uuid.UUID `gorm:"type:uuid;default:null"`
	Deleted     time.Time `gorm:"type:timestamptz;default:null"`
	Active      bool      `gorm:"type:bool;not null"`
	Updated     time.Time `gorm:"type:timestamptz;default:null"`
	UpdatedByID uuid.UUID `gorm:"type:uuid;default:null"`
	Location    string    `gorm:"type:json;default:null"`
	Products    []Product `gorm:"foreignKey:IDCompany;references:UID"`
}
