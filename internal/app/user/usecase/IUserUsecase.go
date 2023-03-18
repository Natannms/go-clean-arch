package usecase

type IUserUsecase interface {
	CreateUser(user *entity.User) error
	GetUserByID(id int64) (*entity.User, error)
}
