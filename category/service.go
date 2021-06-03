package category

import (
	"errors"
	"fmt"
	"ngepet-yuk/entity"
	"ngepet-yuk/helper"
)

type Service interface {
	GetCategories() ([]entity.Category, error)
	SaveNewCategory(input entity.CategoryInput) (entity.Category, error)
	UpdateCategoryByID(ID string, dataInput entity.UpdateCategoryInput) (entity.Category, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCategories() ([]entity.Category, error) {
	categories, err := s.repository.GetAll()

	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (s *service) SaveNewCategory(input entity.CategoryInput) (entity.Category, error) {

	var newCategory = entity.Category{
		Name:        input.Name,
		Description: input.Description,
	}

	createdCategory, err := s.repository.Create(newCategory)

	if err != nil {
		return createdCategory, err
	}

	return createdCategory, nil
}

func (s *service) UpdateCategoryByID(ID string, dataInput entity.UpdateCategoryInput) (entity.Category, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(ID); err != nil {
		return entity.Category{}, err
	}

	userDetail, err := s.repository.FindByID(ID)

	if err != nil {
		return entity.Category{}, err
	}

	if userDetail.ID == 0 {
		newError := fmt.Sprintf("user Detail id %s not found", ID)
		return entity.Category{}, errors.New(newError)
	}

	if dataInput.Name != "" || len(dataInput.Name) != 0 {
		dataUpdate["name"] = dataInput.Name
	}
	if dataInput.Description != "" || len(dataInput.Description) != 0 {
		dataUpdate["description"] = dataInput.Description
	}

	// fmt.Println(dataUpdate)

	cateforyUpdated, err := s.repository.UpdateByID(ID, dataUpdate)

	if err != nil {
		return cateforyUpdated, err
	}

	return cateforyUpdated, nil
}
