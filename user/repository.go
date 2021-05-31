package user

import (
	"ngepet-yuk/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]entity.User, error)
	Create(user entity.User) (entity.User, error)
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
