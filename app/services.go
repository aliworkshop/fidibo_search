package app

import (
	"fidibo/services/fidibo"
)

func (a *App) initFidibo() {
	var err error
	a.fidibo, err = fidibo.NewFidiboClient(a.registry.Sub("fidibo"))
	a.panicOnErr(err)
}
