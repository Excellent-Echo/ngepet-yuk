package entity

import "time"

type User struct {
	ID              int               `gorm:"primaryKey" json:"id"`
	UserName        string            `json:"username"`
	Email           string            `json:"email"`
	Role            string            `json:"role"`
	Password        string            `json:"-"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	UserDetail      UserDetail        `gorm:"foreignKey:UserID"`
	UserTransaction []UserTransaction `gorm:"foreignKey:UserID"`
}

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
