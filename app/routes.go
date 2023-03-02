package app

func (a *App) RegisterRoutes() {
	rg := a.engine.Group("/")
	rg.POST("/search/book", a.MustAuthenticate, a.BookModule.SearchBook)
	rg.POST("/login", a.UserModule.Login)
	rg.POST("/register", a.UserModule.Register)
}
