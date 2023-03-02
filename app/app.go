package app

import (
	"fidibo/book"
	"fidibo/services/fidibo"
	"fidibo/services/redis"
	"fidibo/user"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type App struct {
	config          config
	registry        *viper.Viper
	engine          *gin.Engine
	db              *gorm.DB
	redis           redis.Client
	internalNetwork bool

	fidibo fidibo.Client

	BookModule *book.Module
	UserModule *user.Module
}

func New(registry *viper.Viper) *App {
	return &App{registry: registry}
}

func (a *App) Init() {
	a.initConfig()
	a.initEngine()
	a.initDB()
	a.initRedis()
}

func (a *App) InitModules() {
	a.initBookModule()
	a.initUserModule()
}

func (a *App) InitServices() {
	a.initFidibo()
}

func (a *App) panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
