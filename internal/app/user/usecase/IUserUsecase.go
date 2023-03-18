package usecase

import "github.com/Natannms/go-clean-arch/internal/app/user/entity"

type IUserUsecase interface {
	CreateUser(user *entity.User) error
}
