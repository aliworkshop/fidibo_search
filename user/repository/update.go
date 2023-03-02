package repository

import "fidibo/user/domain"

func (repo *repo) Update(user *domain.User) error {
	err := repo.db.Model(user).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}
