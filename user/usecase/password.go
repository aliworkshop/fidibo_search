package usecase

import (
	"errors"
	"fidibo/user/domain"
	"golang.org/x/crypto/bcrypt"
)

func (uc *usecase) PasswordMatches(password string, user *domain.User) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
