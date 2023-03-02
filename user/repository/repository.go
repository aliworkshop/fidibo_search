package repository

import (
	"fidibo/user/domain"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) domain.UserRepo {
	return &repo{
		db: db,
	}
}
