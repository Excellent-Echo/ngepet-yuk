package entity

import "time"

type UserTransaction struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	SubType      string    `json:"sub_type"`
	Date         time.Time `json:"date"`
	Status       string    `json:"status"`
	TransferFact string    `json:"transfer_fact"`
	UserID       int       `json:"user_id"`
}

type UserTransactionInput struct {
	SubType      string `json:"sub_type" binding:"required"`
	Status       string `json:"status" binding:"required"`
	TransferFact string `json:"transfer_fact" binding:"required"`
}
