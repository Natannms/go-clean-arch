package usecase

import "github.com/Natannms/go-clean-arch/internal/app/user/entity"

type IUserRepository interface {
	Create(user *entity.User) error
	GetByID(id int64) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id int64) error
	GetAll() ([]*entity.User, error)
}

type userUsecase struct {
	userRepo IUserRepository
}

func NewUserUsecase(userRepo IUserRepository) IUserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
