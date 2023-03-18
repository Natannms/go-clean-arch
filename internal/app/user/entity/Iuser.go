package entity

type Iuser interface {
	Create() string
	Get() User
	All() []User
	Update() string
	Delete() string
}
