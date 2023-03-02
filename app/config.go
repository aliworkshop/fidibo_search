package app

import (
	"fidibo/helper"
	"time"
)

const ServiceName = "Fidibo"

type config struct {
	Http struct {
		Address                   string
		Development               bool
		GracefullyShutdownTimeout time.Duration
	}
	Debug bool
	Sql   sqlConfig
	Redis redisConfig
	Jwt   helper.JwtConfig
}

type sqlConfig struct {
	Host               string
	Port               string
	Username           string
	Password           string
	DbName             string
	MaxIdleConnections *int
	MaxOpenConnections *int
	MaxLifetimeSeconds *int
}

type redisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
	Timeout  time.Duration
}

func (c *config) Initialize() {
	if c.Http.GracefullyShutdownTimeout == 0 {
		c.Http.GracefullyShutdownTimeout = time.Second * 10
	}
}
