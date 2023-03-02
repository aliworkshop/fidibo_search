package domain

type UserRepo interface {
	Create(user *User) error
	Update(user *User) error
	GetAll() ([]User, error)
	GetByUsername(username string) (*User, error)
}
