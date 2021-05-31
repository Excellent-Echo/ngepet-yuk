package entity

type UserInput struct {
	UserName string `json:"first_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
