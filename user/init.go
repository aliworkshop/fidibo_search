package user

import (
	"fidibo/user/delivery"
	"fidibo/user/domain"
	"fidibo/user/repository"
	"fidibo/user/usecase"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Module struct {
	UserRepo domain.UserRepo
	UserUc   domain.UserUc

	Login    gin.HandlerFunc
	Register gin.HandlerFunc
}

func New(db *gorm.DB, registry *viper.Viper) *Module {
	m := new(Module)
	m.UserRepo = repository.NewUserRepo(db)
	m.UserUc = usecase.NewUserUc(m.UserRepo, registry)

	m.Login = delivery.NewLoginDelivery(m.UserUc)
	m.Register = delivery.NewRegisterHandler(m.UserUc)
	return m
}
