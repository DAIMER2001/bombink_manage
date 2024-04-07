package models

import (
	"time"

	"github.com/google/uuid"
)

type ProductsStock struct {
	UID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique;not null"`
	Barcodes          string    `gorm:"type:varchar(200);default:null"`
	IDProduct         uuid.UUID `gorm:"type:uuid;not null"`
	Location          string    `gorm:"type:json;default:null"`
	Status            string    `gorm:"type:varchar(200);default:null"`
	SaleDate          time.Time `gorm:"default:null"`
	SalePrice         float64   `gorm:"default:null"`
	ReplacementDate   time.Time `gorm:"default:null"`
	ReplacementReason string    `gorm:"type:text;default:null"`
	ExpirationDate    time.Time `gorm:"default:null"`
	Supplier          string    `gorm:"type:varchar(100);default:null"`
	CommentsSeller    string    `gorm:"type:text;default:null"`
	Product           Product   `gorm:"foreignKey:IDProduct;references:UID"`
}
