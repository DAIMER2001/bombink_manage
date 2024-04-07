package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique;not null"`
	Phone       int64     `gorm:"type:int4;default:null"`
	Password    string    `gorm:"type:varchar(128);not null"`
	LastLogin   time.Time `gorm:"type:timestamptz;default:null"`
	Birthday    time.Time `gorm:"default:null"`
	IsSuperuser bool      `gorm:"type:bool;not null"`
	Username    string    `gorm:"type:varchar(150);not null"`
	FirstName   string    `gorm:"type:varchar(150);not null"`
	LastName    string    `gorm:"type:varchar(150);not null"`
	Email       string    `gorm:"type:varchar(254);not null"`
	IsStaff     bool      `gorm:"type:bool;not null"`
	IsActive    bool      `gorm:"type:bool;not null"`
	DateJoined  time.Time `gorm:"type:timestamptz;not null"`
	Created     time.Time `gorm:"type:timestamptz;default:null"`
	Deleted     time.Time `gorm:"type:timestamptz;default:null"`
	Updated     time.Time `gorm:"type:timestamptz;default:null"`
	CreatedByID uuid.UUID `gorm:"type:uuid;default:null"`
	UpdatedByID uuid.UUID `gorm:"type:uuid;default:null"`
	DeletedByID uuid.UUID `gorm:"type:uuid;default:null"`
	CreatedBy   *User     `gorm:"foreignKey:CreatedByID;references:UID"`
	UpdatedBy   *User     `gorm:"foreignKey:UpdatedByID;references:UID"`
	DeletedBy   *User     `gorm:"foreignKey:DeletedByID;references:UID"`
}
