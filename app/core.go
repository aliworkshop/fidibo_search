package app

import (
	"fidibo/services/redis"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func (a *App) initEngine() {
	a.engine = gin.New()
	if !a.config.Http.Development {
		gin.SetMode(gin.ReleaseMode)
	}
	if os.Getenv("InternalNetwork") == "true" {
		a.internalNetwork = true
	}
}

func (a *App) initConfig() {
	a.panicOnErr(a.registry.Unmarshal(&a.config))
	a.config.Initialize()
}

func (a *App) initDB() {
	if !a.internalNetwork {
		a.config.Sql.Host = "localhost"
	}
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		a.config.Sql.Username,
		a.config.Sql.Password,
		a.config.Sql.Host,
		a.config.Sql.Port,
		a.config.Sql.DbName,
	)
	dialect := mysql.Open(connStr)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             5 * time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn,     // Log level
			IgnoreRecordNotFoundError: false,           // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,            // Disable color
		},
	)
	d, err := gorm.Open(dialect, &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 newLogger,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	a.panicOnErr(err)
	sqlDB, err := d.DB()
	a.panicOnErr(err)
	if a.config.Sql.MaxIdleConnections != nil {
		sqlDB.SetMaxIdleConns(*a.config.Sql.MaxIdleConnections)
	}
	if a.config.Sql.MaxOpenConnections != nil {
		sqlDB.SetMaxOpenConns(*a.config.Sql.MaxOpenConnections)
	}
	if a.config.Sql.MaxLifetimeSeconds != nil {
		sqlDB.SetConnMaxLifetime(time.Second * time.Duration(*a.config.Sql.MaxLifetimeSeconds))
	}
	if a.config.Debug {
		fmt.Println("debug is true")
		d = d.Debug()
	}

	a.db = d
}

func (a *App) initRedis() {
	if !a.internalNetwork {
		a.config.Redis.Host = "localhost"
	}
	a.redis = redis.NewRedisClient(redis.Config{
		Addr:     fmt.Sprintf("%s:%s", a.config.Redis.Host, a.config.Redis.Port),
		Password: a.config.Redis.Password,
		DB:       a.config.Redis.DB,
		Timeout:  a.config.Redis.Timeout,
	})
}
