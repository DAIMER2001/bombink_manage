package dtos

import (
	"time"

	"github.com/google/uuid"
)

type UserDto struct {
	UID         uuid.UUID
	Phone       int64
	Password    string
	LastLogin   time.Time
	Birthday    time.Time
	IsSuperuser bool
	Username    string
	FirstName   string
	LastName    string
	Email       string
	IsStaff     bool
	IsActive    bool
}
