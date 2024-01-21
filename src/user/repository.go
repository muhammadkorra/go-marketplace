package user

type UserRepository interface {
	FindByID(id uint) (*User, error)
	FindAll() ([]User, error)
}
