package category

import (
	"ngepet-yuk/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]entity.Category, error)
	Create(category entity.Category) (entity.Category, error)
	FindByID(ID string) (entity.Category, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]entity.Category, error) {
	var categories []entity.Category

	err := r.db.Find(&categories).Error
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (r *repository) Create(category entity.Category) (entity.Category, error) {

	err := r.db.Create(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Category, error) {

	var category entity.Category

	if err := r.db.Model(&category).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return category, err
	}

	if err := r.db.Where("id = ?", ID).Find(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) FindByID(ID string) (entity.Category, error) {
	var category entity.Category

	err := r.db.Where("id = ?", ID).Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}
