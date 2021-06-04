package userDetail

import (
	"ngepet-yuk/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]entity.UserDetail, error)
	FindByID(userDetailID string) (entity.UserDetail, error)
	FindByUserID(userID string) (entity.UserDetail, error)
	Create(input entity.UserDetail) (entity.UserDetail, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.UserDetail, error)
	// Complete(TodoID string) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]entity.UserDetail, error) {
	var UserDetails []entity.UserDetail

	err := r.db.Find(&UserDetails).Error
	if err != nil {
		return UserDetails, err
	}

	return UserDetails, nil
}

func (r *repository) FindByID(userDetailID string) (entity.UserDetail, error) {
	var userDetail entity.UserDetail

	if err := r.db.Where("id = ?", userDetailID).Find(&userDetail).Error; err != nil {
		return userDetail, err
	}

	return userDetail, nil
}

func (r *repository) FindByUserID(userID string) (entity.UserDetail, error) {
	var userDetail entity.UserDetail

	if err := r.db.Where("user_id = ?", userID).Find(&userDetail).Error; err != nil {
		return userDetail, err
	}

	return userDetail, nil
}

func (r *repository) Create(input entity.UserDetail) (entity.UserDetail, error) {
	if err := r.db.Create(&input).Error; err != nil {
		return input, err
	}

	return input, nil

}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.UserDetail, error) {
	var userDetail entity.UserDetail

	if err := r.db.Model(&userDetail).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return userDetail, err
	}

	if err := r.db.Where("id = ?", ID).Find(&userDetail).Error; err != nil {
		return userDetail, err
	}

	return userDetail, nil
}
