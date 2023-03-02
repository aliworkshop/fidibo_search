package app

import (
	"fidibo/book"
	"fidibo/user"
)

func (a *App) initBookModule() {
	a.BookModule = book.New(a.db, a.fidibo, a.redis)
}

func (a *App) initUserModule() {
	a.UserModule = user.New(a.db, a.registry.Sub("jwt"))
}
