package models

import (
	"time"

	"github.com/google/uuid"
)

type DataList struct {
	DataName   string    `gorm:"type:varchar(100);not null"`
	UID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique;not null"`
	IDTypeList uuid.UUID `gorm:"type:uuid;not null"`
	Data       string    `gorm:"type:json;default:null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`

	// fields for relation
	TypeList TypeList `gorm:"foreignKey:IDTypeList;references:UID"`
}
