package handler

import (
	"ngepet-yuk/auth"
	"ngepet-yuk/entity"
	"ngepet-yuk/helper"
	"ngepet-yuk/userTransaction"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userTransactionHandler struct {
	userTransactionService userTransaction.Service
	authService            auth.Service
}

func NewUserTransactionHandler(userTransactionService userTransaction.Service, authService auth.Service) *userTransactionHandler {
	return &userTransactionHandler{userTransactionService, authService}
}

func (h *userTransactionHandler) GetUserTransactionByUserIDHandler(c *gin.Context) {

	userData := int(c.MustGet("currentUser").(int))

	if userData == 0 {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user not authorize / not login"})

		c.JSON(401, responseError)
		return
	}

	userID := strconv.Itoa(userData)

	userTransaction, err := h.userTransactionService.GetUserTransactionByUserID(userID)

	if err != nil {
		responseError := helper.APIResponse("status unauthorize", 401, "error", gin.H{"error": err.Error()})

		c.JSON(401, responseError)
		return
	}

	response := helper.APIResponse("success get user Transaction by user ID", 200, "success", userTransaction)
	c.JSON(200, response)
}

func (h *userTransactionHandler) SaveNewUserTransactionHandler(c *gin.Context) {
	userData := int(c.MustGet("currentUser").(int))

	if userData == 0 {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user not authorize / not login"})

		c.JSON(401, responseError)
		return
	}

	userID := strconv.Itoa(userData)

	var inputUserTransaction entity.UserTransactionInput

	if err := c.ShouldBindJSON(&inputUserTransaction); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	NewUserTransaction, err := h.userTransactionService.SaveNewUserTransaction(inputUserTransaction, userID)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create new user Transaction", 201, "success", NewUserTransaction)
	c.JSON(201, response)
}
