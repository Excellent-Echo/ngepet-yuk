package userDetail

import (
	"errors"
	"fmt"
	"ngepet-yuk/entity"
	"ngepet-yuk/helper"
	"strconv"
	"time"
)

type Service interface {
	// GetUserDetailByID(userDetailID string) (entity.UserDetail, error)
	GetAllUserDetail() ([]entity.UserDetail, error)
	GetUserDetailByUserID(userID string) (entity.UserDetail, error)
	SaveNewUserDetail(input entity.UserDetailInput, userID string) (entity.UserDetail, error)
	UpdateUserDetailByID(ID string, dataInput entity.UpdateUserDetailInput) (entity.UserDetail, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// func (s *service) GetUserDetailByID(userDetailID string) (entity.UserDetail, error) {
// 	userDetail, err := s.repository.FindByID(userDetailID)

// 	if err != nil {
// 		return userDetail, err
// 	}

// 	return userDetail, nil
// }

func (s *service) GetAllUserDetail() ([]entity.UserDetail, error) {
	userDetail, err := s.repository.GetAll()
	if err != nil {
		return userDetail, err
	}
	return userDetail, nil
}

func (s *service) GetUserDetailByUserID(userID string) (entity.UserDetail, error) {
	userDetail, err := s.repository.FindByUserID(userID)

	if err != nil {
		return userDetail, err
	}

	if userDetail.ID == 0 {
		errStatus := fmt.Sprintf("userdetail for user id %s not created", userID)
		return userDetail, errors.New(errStatus)
	}

	return userDetail, nil
}

func (s *service) SaveNewUserDetail(input entity.UserDetailInput, userID string) (entity.UserDetail, error) {
	IDUser, _ := strconv.Atoi(userID)

	checkStatus, err := s.repository.FindByID(userID)

	if err != nil {
		return checkStatus, err
	}

	if checkStatus.UserID == IDUser {
		errorStatus := fmt.Sprintf("userDetail for user ID : %s has been created", userID)
		return checkStatus, errors.New(errorStatus)
	}

	var NewUserDetail = entity.UserDetail{
		SubType:     input.SubType,
		NoHandphone: input.NoHandphone,
		Gender:      input.Gender,
		Address:     input.Address,
		Period:      input.Period,
		UserID:      IDUser,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createUserDetail, err := s.repository.Create(NewUserDetail)

	if err != nil {
		return createUserDetail, err
	}

	return createUserDetail, nil
}

func (s *service) UpdateUserDetailByID(ID string, dataInput entity.UpdateUserDetailInput) (entity.UserDetail, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(ID); err != nil {
		return entity.UserDetail{}, err
	}

	userDetail, err := s.repository.FindByID(ID)

	if err != nil {
		return entity.UserDetail{}, err
	}

	if userDetail.ID == 0 {
		newError := fmt.Sprintf("user Detail id %s not found", ID)
		return entity.UserDetail{}, errors.New(newError)
	}

	if dataInput.NoHandphone != 0 {
		dataUpdate["no_handphone"] = dataInput.NoHandphone
	}
	if dataInput.Gender != "" || len(dataInput.Gender) != 0 {
		dataUpdate["gender"] = dataInput.Gender
	}
	if dataInput.Address != "" || len(dataInput.Gender) != 0 {
		dataUpdate["address"] = dataInput.Address
	}

	dataUpdate["updated_at"] = time.Now()

	// fmt.Println(dataUpdate)

	userDetailUpdate, err := s.repository.UpdateByID(ID, dataUpdate)

	if err != nil {
		return userDetailUpdate, err
	}

	return userDetailUpdate, nil
}
