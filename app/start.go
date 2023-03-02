package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func (a *App) Start() {
	a.RegisterRoutes()
	go func() {
		if err := a.engine.Run(a.config.Http.Address); err != nil {
			panic(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	<-c
	a.Stop()
}

func (a *App) Stop() {
	log.Println("shutting down...")
}
