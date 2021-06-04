package mastery

import (
	"ngepet-yuk/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]entity.Mastery, error)
	Create(mastery entity.Mastery) (entity.Mastery, error)
	FindByID(ID string) (entity.Mastery, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Mastery, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]entity.Mastery, error) {
	var masteries []entity.Mastery

	err := r.db.Find(&masteries).Error
	if err != nil {
		return masteries, err
	}

	return masteries, nil
}

func (r *repository) Create(mastery entity.Mastery) (entity.Mastery, error) {

	err := r.db.Create(&mastery).Error
	if err != nil {
		return mastery, err
	}

	return mastery, nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Mastery, error) {

	var mastery entity.Mastery

	if err := r.db.Model(&mastery).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return mastery, err
	}

	if err := r.db.Where("id = ?", ID).Find(&mastery).Error; err != nil {
		return mastery, err
	}

	return mastery, nil
}

func (r *repository) FindByID(ID string) (entity.Mastery, error) {
	var mastery entity.Mastery

	err := r.db.Where("id = ?", ID).Find(&mastery).Error
	if err != nil {
		return mastery, err
	}

	return mastery, nil
}
