package repositories

import (
	"github.com/daimer/bombink-manage/internal/domain/user/entities"
	"github.com/google/uuid"
)

type IUserRepository interface {
	SaveUser(*entities.User) (uuid.UUID, error)
}
