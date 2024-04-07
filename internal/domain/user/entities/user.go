package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UID         uuid.UUID
	Phone       int64
	Password    string
	Birthday    time.Time
	LastLogin   time.Time
	IsSuperuser bool
	Username    string
	FirstName   string
	LastName    string
	Email       string
	IsStaff     bool
	IsActive    bool
	DateJoined  time.Time
	Created     time.Time
	Deleted     time.Time
	Updated     time.Time
	CreatedByID uuid.UUID
	UpdatedByID uuid.UUID
	DeletedByID uuid.UUID
	CreatedBy   *User
	UpdatedBy   *User
	DeletedBy   *User
}
