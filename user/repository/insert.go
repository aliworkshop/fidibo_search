package repository

import (
	"fidibo/user/domain"
	"github.com/go-sql-driver/mysql"
)

func (repo *repo) Create(user *domain.User) error {
	err := repo.db.Model(user).Create(user).Error
	if err != nil {
		if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
			return nil
		}
		return err
	}
	return nil
}
