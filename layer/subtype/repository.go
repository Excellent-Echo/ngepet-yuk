package subtype

import (
	"ngepet-yuk/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]entity.SubType, error)
	Create(subtype entity.SubType) (entity.SubType, error)
	FindByID(ID string) (entity.SubType, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.SubType, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]entity.SubType, error) {
	var SubTypes []entity.SubType

	err := r.db.Find(&SubTypes).Error
	if err != nil {
		return SubTypes, err
	}

	return SubTypes, nil
}

func (r *repository) Create(subtype entity.SubType) (entity.SubType, error) {

	err := r.db.Create(&subtype).Error
	if err != nil {
		return subtype, err
	}

	return subtype, nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.SubType, error) {

	var subtype entity.SubType

	if err := r.db.Model(&subtype).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return subtype, err
	}

	if err := r.db.Where("id = ?", ID).Find(&subtype).Error; err != nil {
		return subtype, err
	}

	return subtype, nil
}

func (r *repository) FindByID(ID string) (entity.SubType, error) {
	var subtype entity.SubType

	err := r.db.Where("id = ?", ID).Find(&subtype).Error
	if err != nil {
		return subtype, err
	}

	return subtype, nil
}
