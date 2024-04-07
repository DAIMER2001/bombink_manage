package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text;not null"`
	Price       int64     `gorm:"type:int4;not null"`
	Discount    int64     `gorm:"type:int4;default:null"`
	Created     time.Time `gorm:"type:timestamptz;not null"`
	Deleted     time.Time `gorm:"type:timestamptz;default:null"`
	Updated     time.Time `gorm:"type:timestamptz;default:null"`
	UID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique;not null"`
	CreatedByID uuid.UUID `gorm:"type:uuid;default:null"`
	UpdatedByID uuid.UUID `gorm:"type:uuid;default:null"`
	Active      bool      `gorm:"type:bool;not null"`
	DeletedByID uuid.UUID `gorm:"type:uuid;default:null"`
	Category    string    `gorm:"type:varchar(100);default:null"`
	Stock       int64     `gorm:"type:int8;default:null"`
	IDCompany   uuid.UUID `gorm:"type:uuid;not null"`
	CreatedBy   *User     `gorm:"foreignKey:CreatedByID;references:UID"`
	UpdatedBy   *User     `gorm:"foreignKey:UpdatedByID;references:UID"`
	DeletedBy   *User     `gorm:"foreignKey:DeletedByID;references:UID"`
}
