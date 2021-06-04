package mastery

import (
	"errors"
	"fmt"
	"ngepet-yuk/entity"
	"ngepet-yuk/helper"
)

type Service interface {
	GetMasteries() ([]entity.Mastery, error)
	SaveNewMastery(mastery entity.MasteryInput) (entity.Mastery, error)
	UpdateMasteryByID(ID string, dataInput entity.UpdateMasteryInput) (entity.Mastery, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetMasteries() ([]entity.Mastery, error) {
	masteries, err := s.repository.GetAll()

	if err != nil {
		return masteries, err
	}

	return masteries, nil
}

func (s *service) SaveNewMastery(mastery entity.MasteryInput) (entity.Mastery, error) {

	var newMastery = entity.Mastery{
		Level: mastery.Level,
	}

	createdMastery, err := s.repository.Create(newMastery)

	if err != nil {
		return createdMastery, err
	}

	return createdMastery, nil
}

func (s *service) UpdateMasteryByID(ID string, dataInput entity.UpdateMasteryInput) (entity.Mastery, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(ID); err != nil {
		return entity.Mastery{}, err
	}

	mastery, err := s.repository.FindByID(ID)

	if err != nil {
		return entity.Mastery{}, err
	}

	if mastery.ID == 0 {
		newError := fmt.Sprintf("mastery id %s not found", ID)
		return entity.Mastery{}, errors.New(newError)
	}

	if dataInput.Level != "" || len(dataInput.Level) != 0 {
		dataUpdate["level"] = dataInput.Level
	}

	// fmt.Println(dataUpdate)

	masteryUpdated, err := s.repository.UpdateByID(ID, dataUpdate)

	if err != nil {
		return masteryUpdated, err
	}

	return masteryUpdated, nil
}
