package usecase

import (
	"fidibo/book/domain"
	"fidibo/services/fidibo"
	"fidibo/services/redis"
)

type usecase struct {
	repo   domain.BookRepo
	fidibo fidibo.Client
	redis  redis.Client
}

func NewBookUc(db domain.BookRepo, fidibo fidibo.Client, redis redis.Client) domain.BookUc {
	return &usecase{
		repo:   db,
		fidibo: fidibo,
		redis:  redis,
	}
}
