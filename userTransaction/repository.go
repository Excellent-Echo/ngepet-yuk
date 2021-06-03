package userTransaction

import (
	"ngepet-yuk/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindByID(userTransactionID string) (entity.UserTransaction, error)
	FindByUserID(userID string) (entity.UserTransaction, error)
	Create(input entity.UserTransaction) (entity.UserTransaction, error)
	// Complete(TodoID string) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByID(userTransactionID string) (entity.UserTransaction, error) {
	var userTransaction entity.UserTransaction

	if err := r.db.Where("id = ?", userTransactionID).Find(&userTransaction).Error; err != nil {
		return userTransaction, err
	}

	return userTransaction, nil
}

func (r *repository) FindByUserTransactionID(userID string) (entity.UserTransaction, error) {
	var userTransaction entity.UserTransaction

	if err := r.db.Where("user_id = ?", userID).Find(&userTransaction).Error; err != nil {
		return userTransaction, err
	}

	return userTransaction, nil
}

func (r *repository) Create(input entity.UserTransaction) (entity.UserTransaction, error) {
	if err := r.db.Create(&input).Error; err != nil {
		return input, err
	}

	return input, nil

}
