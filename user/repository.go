package user

import (
	"ngepet-yuk/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]entity.User, error)
	Create(user entity.User) (entity.User, error)
	FindByID(ID string) (entity.User, error)
	DeleteByID(ID string) (string, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]entity.User, error) {
	var Users []entity.User

	err := r.db.Find(&Users).Error
	if err != nil {
		return Users, err
	}

	return Users, nil
}

func (r *repository) Create(user entity.User) (entity.User, error) {

	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(ID string) (entity.User, error) {
	var user entity.User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) DeleteByID(ID string) (string, error) {
	if err := r.db.Where("id = ?", ID).Delete(&entity.User{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.User, error) {

	var user entity.User

	if err := r.db.Model(&user).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return user, err
	}

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
