package usecase

import "fidibo/user/domain"

func (uc *usecase) Create(user *domain.User) error {
	return uc.repo.Create(user)
}
