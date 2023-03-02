package domain

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type RegisterRequest struct {
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	Email     string `json:"email" form:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
