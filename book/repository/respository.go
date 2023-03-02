package repository

import (
	"fidibo/book/domain"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) domain.BookRepo {
	return &repo{
		db: db,
	}
}
