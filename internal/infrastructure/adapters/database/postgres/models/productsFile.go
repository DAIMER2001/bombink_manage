package models

import (
	"github.com/google/uuid"
)

type ProductsFile struct {
	MediaURL  string    `gorm:"type:varchar(150);default:null"`
	IDProduct uuid.UUID `gorm:"type:uuid;not null"`
	UID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique;not null"`
	TypeMedia string    `gorm:"type:varchar(30);default:null"`
	Product   Product   `gorm:"foreignKey:IDProduct;references:UID"`
}
