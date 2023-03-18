package usecase

type IUserRepository interface {
}

type userUsecase struct {
	userRepo IUserRepository
}

func NewUserUsecase(userRepo IUserRepository) IUserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
