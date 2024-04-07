package User

import (
	"errors"

	"github.com/daimer/bombink-manage/internal/application/usecases/saveUser/dtos"
	"github.com/daimer/bombink-manage/internal/domain/user/entities"
	"github.com/daimer/bombink-manage/internal/domain/user/repositories"
	"github.com/daimer/bombink-manage/internal/domain/utils/logger"
	"github.com/google/uuid"
)

func (a *UserUsecase) SaveUser(requestBody *dtos.UserDto) (string, error) {
	/*==================================================
	***************** VALID BODY User ***************
	==================================================*/
	err := validBodyUser(requestBody)
	if err != nil {
		return "", err
	}
	/*==================================================
	************ CREATE_ACTIVE_User *****************
	==================================================*/
	idUser, err := createUser(
		a.UserRepo,
		requestBody)

	if err != nil {
		return "", errors.New("UserNotFound")
	}
	return idUser.String(), nil
}

var createUser = func(
	UserRepo repositories.IUserRepository,
	requestBody *dtos.UserDto,
) (uuid.UUID, error) {
	/*==================================================
	************ CREATE_User *****************
	==================================================*/

	UserEntity := &entities.User{
		UID:         requestBody.UID,
		Phone:       requestBody.Phone,
		Password:    requestBody.Password,
		Birthday:    requestBody.Birthday,
		IsSuperuser: requestBody.IsSuperuser,
		Username:    requestBody.Username,
		FirstName:   requestBody.FirstName,
		LastName:    requestBody.LastName,
		Email:       requestBody.Email,
		IsStaff:     requestBody.IsStaff,
		IsActive:    requestBody.IsActive,
	}
	idUser, err := UserRepo.SaveUser(UserEntity)
	if err != nil {
		logger.Error("error saving User", err)
		return uuid.Nil, errors.New("SaveActiveUserFailed")
	}
	return idUser, err
}

var validBodyUser = func(
	requestBody *dtos.UserDto,
) error {
	if len(requestBody.Username) > 100 {
		return errors.New("ErrorLenName")
	}
	if len(requestBody.FirstName) > 100 {
		return errors.New("ErrorLenName")
	}
	if len(requestBody.LastName) > 100 {
		return errors.New("ErrorLenName")
	}
	if len(requestBody.Email) > 150 {
		return errors.New("ErrorLenName")
	}
	if len(requestBody.Password) > 5 {
		return errors.New("ErrorLenVersion")
	}
	return nil
}
