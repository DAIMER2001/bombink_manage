package database

import (
	"github.com/daimer/bombink-manage/internal/domain/user/entities"
	"github.com/daimer/bombink-manage/internal/infrastructure/adapters/database/postgres/models"
	"github.com/google/uuid"
)

func (adapter *BombinkDB) SaveUser(userEntity *entities.User) (uuid.UUID, error) {
	var modelSave = models.User{
		UID:         userEntity.UID,
		Phone:       userEntity.Phone,
		Password:    userEntity.Password,
		LastLogin:   userEntity.LastLogin,
		IsSuperuser: userEntity.IsSuperuser,
		Username:    userEntity.Username,
		FirstName:   userEntity.FirstName,
		LastName:    userEntity.LastName,
		Email:       userEntity.Email,
		IsStaff:     userEntity.IsStaff,
		IsActive:    userEntity.IsActive,
	}
	result := adapter.conn.Debug().Save(&modelSave)
	if result.Error != nil {
		return uuid.UUID{}, result.Error
	}
	return modelSave.UID, nil
}
