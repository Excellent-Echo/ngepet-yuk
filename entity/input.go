package entity

type LoginUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserInput struct {
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	UsernName string `json:"username"`
	Email     string `json:"email"`
}
