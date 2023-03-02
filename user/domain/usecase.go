package domain

type UserUc interface {
	Create(user *User) error
	Update(user *User) error
	GetAll() ([]User, error)
	GetByUsername(username string) (*User, error)
	PasswordMatches(password string, user *User) (bool, error)
	Login(user *User) (string, int, error)
}
