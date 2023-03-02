package usecase

import (
	"fidibo/helper"
	"fidibo/user/domain"
	"github.com/spf13/viper"
)

type usecase struct {
	repo   domain.UserRepo
	config helper.JwtConfig
}

func (uc *usecase) GetAll() ([]domain.User, error) {
	return uc.repo.GetAll()
}

func (uc *usecase) GetByUsername(username string) (*domain.User, error) {
	return uc.repo.GetByUsername(username)
}

func NewUserUc(db domain.UserRepo, registry *viper.Viper) domain.UserUc {
	uc := &usecase{
		repo: db,
	}
	if err := registry.Unmarshal(&uc.config); err != nil {
		panic(err)
	}

	return uc
}
