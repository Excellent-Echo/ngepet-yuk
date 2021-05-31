package entity

import "time"

type UserTransaction struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	SubType      int       `json:"sub_type"`
	Date         time.Time `json:"date"`
	Status       string    `json:"status"`
	TransferFact string    `json:"transfer_fact"`
	UserID       int       `json:"user_id"`
}
