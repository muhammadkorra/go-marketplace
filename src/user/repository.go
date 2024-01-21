package user

type UserRepository interface {
	FindById(id uint) (*User, error)
	FindAll() ([]User, error)
}
