package repository

import "fidibo/user/domain"

func (repo *repo) GetAll() ([]domain.User, error) {
	users := make([]domain.User, 0)
	err := repo.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *repo) GetByUsername(username string) (user *domain.User, err error) {
	err = repo.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return
	}
	return
}
