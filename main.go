package main

import (
	"fidibo/app"
	"github.com/spf13/viper"
	"os"
)

func main() {
	fidibo := app.New(config())
	fidibo.Init()
	fidibo.InitServices()
	fidibo.InitModules()

	fidibo.Start()
}

func config() *viper.Viper {
	v := viper.New()
	v.SetConfigType("yaml")
	f, err := os.Open("./config.yaml")
	if err != nil {
		panic("cannot read config: " + err.Error())
	}
	err = v.ReadConfig(f)
	if err != nil {
		panic("cannot read config" + err.Error())
	}

	return v
}
