package User

import (
	"github.com/daimer/bombink-manage/internal/domain/user/repositories"
)

type UserUsecase struct {
	UserRepo repositories.IUserRepository
}

func New(
	UserRepo repositories.IUserRepository,
) *UserUsecase {
	return &UserUsecase{
		UserRepo: UserRepo,
	}
}
