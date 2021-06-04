package subtype

import (
	"errors"
	"fmt"
	"ngepet-yuk/entity"
	"ngepet-yuk/helper"
)

type Service interface {
	GetAllSubtype() ([]SubtypeFormat, error)
	SaveNewSubtype(subtype entity.SubTypeInput) (SubtypeFormat, error)
	UpdateSubTypeByID(subtypeID string, dataInput entity.UpdateSubTypeInput) (SubtypeFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllSubtype() ([]SubtypeFormat, error) {
	Subtypes, err := s.repository.GetAll()
	var formatSubtypes []SubtypeFormat

	for _, subtype := range Subtypes {
		formatSubtype := FormatSubtype(subtype)
		formatSubtypes = append(formatSubtypes, formatSubtype)

	}
	if err != nil {
		return formatSubtypes, err
	}

	return formatSubtypes, nil
}

func (s *service) SaveNewSubtype(subtype entity.SubTypeInput) (SubtypeFormat, error) {

	var newSubtype = entity.SubType{
		SubType: subtype.SubType,
		Price:   subtype.Price,
		Period:  subtype.Period,
	}

	createSubtype, err := s.repository.Create(newSubtype)
	formatSubtype := FormatSubtype(createSubtype)

	if err != nil {
		return formatSubtype, err
	}

	return formatSubtype, nil
}

func (s *service) UpdateSubTypeByID(subtypeID string, dataInput entity.UpdateSubTypeInput) (SubtypeFormat, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(subtypeID); err != nil {
		return SubtypeFormat{}, err
	}

	subtype, err := s.repository.FindByID(subtypeID)

	if err != nil {
		return SubtypeFormat{}, err
	}

	if subtype.ID == 0 {
		newError := fmt.Sprintf("subtype id %s not found", subtypeID)
		return SubtypeFormat{}, errors.New(newError)
	}

	if dataInput.SubType != "" || len(dataInput.SubType) != 0 {
		dataUpdate["sub_type"] = dataInput.SubType
	}
	if dataInput.Price >= 0 {
		dataUpdate["price"] = dataInput.Price
	}

	if dataInput.Price >= 0 {
		dataUpdate["period"] = dataInput.Price
	}

	// fmt.Println(dataUpdate)

	subtypeUpdated, err := s.repository.UpdateByID(subtypeID, dataUpdate)

	if err != nil {
		return SubtypeFormat{}, err
	}

	formatSubtype := FormatSubtype(subtypeUpdated)

	return formatSubtype, nil
}
