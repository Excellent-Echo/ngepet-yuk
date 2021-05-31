package entity

import "time"

type User struct {
	ID              int               `gorm:"primaryKey" json:"id"`
	UserName        string            `gorm:"uniqueIndex" json:"username"`
	Email           string            `gorm:"uniqueIndex" json:"email"`
	Role            string            `json:"role"`
	Password        string            `json:"-"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	UserDetail      UserDetail        `gorm:"foreignKey:UserID"`
	UserTransaction []UserTransaction `gorm:"foreignKey:UserID"`
}
