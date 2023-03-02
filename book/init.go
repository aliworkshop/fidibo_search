package book

import (
	"fidibo/book/delivery"
	"fidibo/book/domain"
	"fidibo/book/repository"
	"fidibo/book/usecase"
	"fidibo/services/fidibo"
	"fidibo/services/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	BookRepo domain.BookRepo
	BookUc   domain.BookUc

	SearchBook gin.HandlerFunc
}

func New(db *gorm.DB, fidibo fidibo.Client, redis redis.Client) *Module {
	m := new(Module)
	m.BookRepo = repository.NewBookRepo(db)
	m.BookUc = usecase.NewBookUc(m.BookRepo, fidibo, redis)

	m.SearchBook = delivery.SearchBookDelivery(m.BookUc)
	return m
}
