package usecase

import "fidibo/user/domain"

func (uc *usecase) Update(user *domain.User) error {
	return uc.repo.Update(user)
}
