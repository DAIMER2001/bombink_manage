package models

import (
	"time"

	"github.com/google/uuid"
)

type Sale struct {
	Amount      float64   `gorm:"type:numeric(10, 2);not null"`
	SaleDate    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;null"`
	UID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique;not null"`
	IDSeller    uuid.UUID `gorm:"type:uuid;not null"`
	IDBuyer     uuid.UUID `gorm:"type:uuid;not null"`
	Created     time.Time `gorm:"type:timestamptz;not null"`
	Deleted     time.Time `gorm:"type:timestamptz;default:null"`
	Updated     time.Time `gorm:"type:timestamptz;default:null"`
	CreatedByID uuid.UUID `gorm:"type:uuid;default:null"`
	UpdatedByID uuid.UUID `gorm:"type:uuid;default:null"`
	DeletedByID uuid.UUID `gorm:"type:uuid;default:null"`
	Seller      Seller    `gorm:"foreignKey:IDSeller;references:UID"`
	Buyer       User      `gorm:"foreignKey:IDBuyer;references:UID"`
}
