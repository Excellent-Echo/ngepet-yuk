package entity

import "time"

type UserDetail struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	SubType     string    `json:"sub_type"`
	NoHandphone int       `json:"no_handphone"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	Period      int       `json:"period"`
	UserID      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
