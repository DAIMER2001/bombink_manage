package models

import (
	"time"

	"github.com/google/uuid"
)

type TattooReactions struct {
	UID                uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique;not null"`
	ReactionType       string            `gorm:"type:varchar(40);default:null"`
	Reaction           string            `gorm:"type:varchar(40);default:null"`
	Comments           string            `gorm:"type:text;default:null"`
	IDTattoo           uuid.UUID         `gorm:"type:uuid;not null"`
	IDResponseReaction uuid.UUID         `gorm:"type:uuid;default:null"`
	Created            time.Time         `gorm:"type:timestamptz;not null"`
	CreatedByID        uuid.UUID         `gorm:"type:uuid;default:null"`
	Deleted            time.Time         `gorm:"type:timestamptz;default:null"`
	Updated            time.Time         `gorm:"type:timestamptz;default:null"`
	UpdatedByID        uuid.UUID         `gorm:"type:uuid;default:null"`
	DeletedByID        uuid.UUID         `gorm:"type:uuid;default:null"`
	Tattoo             Tattoos           `gorm:"foreignKey:IDTattoo;references:UID"`
	ParentID           *uuid.UUID        `gorm:"type:uuid;default:null"`
	ResponseReaction   []TattooReactions `gorm:"foreignKey:ParentID;references:UID"`
	Parent             *TattooReactions  `gorm:"foreignKey:ParentID;references:UID"`
}
