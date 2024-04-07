package models

import (
	"time"

	"github.com/google/uuid"
)

type Dropshipping struct {
	ProductName   string    `gorm:"type:varchar(100);not null"`
	Quantity      int64     `gorm:"not null"`
	Price         float64   `gorm:"type:numeric(10, 2);not null"`
	OrderDate     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;null"`
	UID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique;not null"`
	IDDropshipper uuid.UUID `gorm:"type:uuid;not null"`
	IDProduct     uuid.UUID `gorm:"type:uuid;not null"`
	Dropshipper   Seller    `gorm:"foreignKey:IDDropshipper;references:UID"`
	Product       Product   `gorm:"foreignKey:IDProduct;references:UID"`
}
