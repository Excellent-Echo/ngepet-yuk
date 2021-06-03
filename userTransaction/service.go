package userTransaction

import (
	"errors"
	"fmt"
	"ngepet-yuk/entity"
	"strconv"
	"time"
)

type Service interface {
	GetUserTransactionByUserID(userID string) (entity.UserTransaction, error)
	SaveNewUserTransaction(input entity.UserTransactionInput, userID string) (entity.UserTransaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetUserTransactionByUserID(userID string) (entity.UserTransaction, error) {
	userTransaction, err := s.repository.FindByUserTransactionID(userID)

	if err != nil {
		return userTransaction, err
	}

	if userTransaction.ID == 0 {
		errStatus := fmt.Sprintf("userTransaction for user id %s not created", userID)
		return userTransaction, errors.New(errStatus)
	}

	return userTransaction, nil
}

func (s *service) SaveNewUserTransaction(input entity.UserTransactionInput, userID string) (entity.UserTransaction, error) {
	IDUser, _ := strconv.Atoi(userID)

	checkStatus, err := s.repository.FindByID(userID)

	if err != nil {
		return checkStatus, err
	}

	if checkStatus.UserID == IDUser {
		errorStatus := fmt.Sprintf("userTransaction for user ID : %s has been created", userID)
		return checkStatus, errors.New(errorStatus)
	}

	var NewUserTransaction = entity.UserTransaction{
		SubType:      input.SubType,
		Status:       input.Status,
		TransferFact: input.TransferFact,
		Date:         time.Now(),
	}

	createUserTransaction, err := s.repository.Create(NewUserTransaction)

	if err != nil {
		return createUserTransaction, err
	}

	return createUserTransaction, nil
}
