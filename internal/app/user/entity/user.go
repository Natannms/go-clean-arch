package entity

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func (u *User) Create(user User) string {
	u.Name = user.Name
	u.Email = user.Email
	u.Password = user.Password
	return "Usuário Criado com sucesso!"
}
func (u *User) Get(User) User {
	return User{}
}
func (u *User) All() []User {
	return []User{}
}
func (u *User) Update() string {
	return "Usuário atualizado com sucesso!"
}
func (u *User) Delete() string {
	return "Usuário deletado com sucesso!"
}
