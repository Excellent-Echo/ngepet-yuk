package user

import (
	"errors"
	"fmt"
	"ngepet-yuk/entity"
	"ngepet-yuk/helper"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetAllUser() ([]UserFormat, error)
	SaveNewUser(user entity.UserInput) (UserFormat, error)
	GetUserByID(userID string) (UserFormat, error)
	DeleteUserByID(userID string) (interface{}, error)
	UpdateUserByID(userID string, dataInput entity.UpdateUserInput) (UserFormat, error)
	LoginUser(input entity.LoginUserInput) (entity.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) LoginUser(input entity.LoginUserInput) (entity.User, error) {
	user, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %v not found", user.ID)
		return user, errors.New(newError)
	}

	//ngecek password sama atau tidak
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return user, errors.New("password invalid")
	}

	return user, nil
}

func (s *service) GetAllUser() ([]UserFormat, error) {
	Users, err := s.repository.GetAll()
	var formatUsers []UserFormat

	for _, user := range Users {
		formatUser := FormatUser(user)
		formatUsers = append(formatUsers, formatUser)

	}
	if err != nil {
		return formatUsers, err
	}

	return formatUsers, nil
}

func (s *service) SaveNewUser(user entity.UserInput) (UserFormat, error) {
	//bcryp password jadi hash agar susah diretas
	genPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
		return UserFormat{}, err
	}

	var newUser = entity.User{
		UserName:  user.UserName,
		Email:     user.Email,
		Password:  string(genPassword), //diganti karena untuk hasing
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createUser, err := s.repository.Create(newUser)
	formatUser := FormatUser(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil
}

func (s *service) GetUserByID(userID string) (UserFormat, error) {
	user, err := s.repository.FindByID(userID)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return UserFormat{}, errors.New(newError)
	}

	formatUser := FormatUser(user)

	return formatUser, nil
}

func (s *service) DeleteUserByID(userID string) (interface{}, error) {
	if err := helper.ValidateIDNumber(userID); err != nil {
		return nil, err
	}

	user, err := s.repository.FindByID(userID)

	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteByID(userID)

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("success delete user ID : %s", userID)

	formatDelete := FormatDeleteUser(msg)

	return formatDelete, nil
}

//(UserFormat, error)

func (s *service) UpdateUserByID(userID string, dataInput entity.UpdateUserInput) (UserFormat, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(userID); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.FindByID(userID)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s not found", userID)
		return UserFormat{}, errors.New(newError)
	}

	if dataInput.UsernName != "" || len(dataInput.UsernName) != 0 {
		dataUpdate["username"] = dataInput.UsernName
	}
	if dataInput.Email != "" || len(dataInput.Email) != 0 {
		dataUpdate["email"] = dataInput.Email
	}

	dataUpdate["updated_at"] = time.Now()

	// fmt.Println(dataUpdate)

	userUpdated, err := s.repository.UpdateByID(userID, dataUpdate)

	if err != nil {
		return UserFormat{}, err
	}

	formatUser := FormatUser(userUpdated)

	return formatUser, nil
}
