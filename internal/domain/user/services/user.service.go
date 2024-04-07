package services

import (
	"github.com/daimer/bombink-manage/internal/domain/user/entities"
)

type IUserService interface {
	GetUsersByIDCompanyByName(status, order, name string, idCompany, pageSize, offset int) ([]*entities.User, error)
	GetSearchUserByName(status, order, search string, pageSize, offset int) ([]*entities.User, error)
	GetUserByIDUser(status, order, name string, pageSize, offset int) ([]*entities.User, error)
}
