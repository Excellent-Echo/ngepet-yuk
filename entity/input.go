package entity

type UserInput struct {
	First_name string `json:"first_name" binding:"required"`
	Last_name  string `json:"last_name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}
