package course

import (
	"ngepet-yuk/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]entity.Courses, error)
	Create(course entity.Courses) (entity.Courses, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Courses, error)
	DeleteByID(ID string) (string, error)
	FindByID(ID string) (entity.Courses, error)
	GetByCategoryID(ID string) ([]entity.Courses, error)
	GetByMasteryID(ID string) ([]entity.Courses, error)
	GetBySubcription(ID string) ([]entity.Courses, error)
	GetByCategoryAndMasteryAndSubcription(CID string, MID string, SID string) ([]entity.Courses, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]entity.Courses, error) {
	var courses []entity.Courses

	err := r.db.Find(&courses).Error
	if err != nil {
		return courses, err
	}

	return courses, nil
}

func (r *repository) Create(course entity.Courses) (entity.Courses, error) {

	err := r.db.Create(&course).Error
	if err != nil {
		return course, err
	}

	return course, nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Courses, error) {

	var course entity.Courses

	if err := r.db.Model(&course).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return course, err
	}

	if err := r.db.Where("id = ?", ID).Find(&course).Error; err != nil {
		return course, err
	}

	return course, nil
}

func (r *repository) DeleteByID(ID string) (string, error) {
	if err := r.db.Where("id = ?", ID).Delete(&entity.Courses{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}

func (r *repository) FindByID(ID string) (entity.Courses, error) {
	var course entity.Courses

	err := r.db.Where("id = ?", ID).Find(&course).Error
	if err != nil {
		return course, err
	}

	return course, nil
}

func (r *repository) GetByCategoryID(ID string) ([]entity.Courses, error) {
	var courses []entity.Courses

	err := r.db.Where("category_id = ?", ID).Find(&courses).Error
	if err != nil {
		return courses, err
	}

	return courses, nil
}

func (r *repository) GetByMasteryID(ID string) ([]entity.Courses, error) {
	var courses []entity.Courses

	err := r.db.Where("mastery_id = ?", ID).Find(&courses).Error
	if err != nil {
		return courses, err
	}

	return courses, nil
}

func (r *repository) GetBySubcription(ID string) ([]entity.Courses, error) {
	var courses []entity.Courses

	err := r.db.Where("sub_type = ?", ID).Find(&courses).Error
	if err != nil {
		return courses, err
	}

	return courses, nil
}

func (r *repository) GetByCategoryAndMasteryAndSubcription(CID string, MID string, SID string) ([]entity.Courses, error) {
	var courses []entity.Courses

	err := r.db.Where("category_id = ? AND mastery_id = ? AND sub_type = ?", CID, MID, SID).Find(&courses).Error
	if err != nil {
		return courses, err
	}

	return courses, nil
}
