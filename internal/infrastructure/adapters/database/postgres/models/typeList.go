package models

import (
	"time"

	"github.com/google/uuid"
)

type TypeList struct {
	TypeName  string    `gorm:"type:varchar(100);not null"`
	UID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique;not null"`
	Desc      string    `gorm:"type:text;default:null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
