package handler

import (
	"net/http"
	"strconv"

	"ngepet-yuk/auth"
	"ngepet-yuk/config"
	"ngepet-yuk/entity"
	"ngepet-yuk/helper"
	"ngepet-yuk/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

// SHOW ALL USER
func (h *userHandler) ShowAllUser(c *gin.Context) {
	users, err := h.userService.GetAllUser()

	if err != nil {
		responseError := helper.APIResponse("internal server error", http.StatusOK, "error", gin.H{"errors": err.Error()})

		c.JSON(http.StatusOK, responseError)
		return
	}

	response := helper.APIResponse("success get all user", http.StatusOK, "status OK", users)
	c.JSON(http.StatusOK, response)
}

// CREATE NEW USER OR REGISTER
func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var inputUser entity.UserInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {

		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", http.StatusOK, "bad request", gin.H{"errors": splitError})

		c.JSON(http.StatusOK, responseError)
		return
	}

	// max len password 6
	if err := helper.ValidatePassword(inputUser.Password); err != nil {

		responseError := helper.APIResponse("input data required", http.StatusOK, "bad request", gin.H{"error": "error validate password length < 6"})

		c.JSON(http.StatusOK, responseError)
		return
	}

	newUser, err := h.userService.SaveNewUser(inputUser)
	if err != nil {
		responseError := helper.APIResponse("internal server error", http.StatusOK, "error", gin.H{"errors": err.Error()})

		c.JSON(http.StatusOK, responseError)
		return
	}
	response := helper.APIResponse("success create new User", http.StatusOK, "Status OK", newUser)
	c.JSON(http.StatusOK, response)
}

// GET USER BY ID
func (h *userHandler) GetUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		responseError := helper.APIResponse("input params error", http.StatusOK, "bad request", gin.H{"errors": err.Error()})

		c.JSON(http.StatusOK, responseError)
		return
	}

	response := helper.APIResponse("success get user by ID", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, response)
}

var DB = config.Connection()

// DELETE USER BY ID
func (h *userHandler) DeleteUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	user, err := h.userService.DeleteUserByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request delete user", http.StatusOK, "error", gin.H{"error": err.Error()})

		c.JSON(http.StatusOK, responseError)
		return
	}

	response := helper.APIResponse("success delete user by ID", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, response)
}

// UPDATE USER BY ID
func (h *userHandler) UpdateUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	var updateUserInput entity.UpdateUserInput

	if err := c.ShouldBindJSON(&updateUserInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", http.StatusOK, "bad request", gin.H{"errors": splitError})

		c.JSON(http.StatusOK, responseError)
		return
	}

	idParam, _ := strconv.Atoi(id)

	// ngecek id sama apa engga sama yang di inputin
	userData := int(c.MustGet("currentUser").(int))

	if idParam != userData {
		responseError := helper.APIResponse("Unauthorize", http.StatusUnauthorized, "error", gin.H{"error": "user ID not authorize"})

		c.JSON(http.StatusUnauthorized, responseError)
		return
	}

	user, err := h.userService.UpdateUserByID(id, updateUserInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update user by ID", http.StatusOK, "success", user)
	c.JSON(http.StatusOK, response)
}

// USER LOGIN
func (h *userHandler) LoginUserHandler(c *gin.Context) {
	var inputLoginUser entity.LoginUserInput

	if err := c.ShouldBindJSON(&inputLoginUser); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", http.StatusOK, "bad request", gin.H{"errors": splitError})

		c.JSON(http.StatusOK, responseError)
		return
	}

	userData, err := h.userService.LoginUser(inputLoginUser)

	if err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", http.StatusUnauthorized, "bad request", gin.H{"errors": splitError})

		c.JSON(http.StatusUnauthorized, responseError)
		return
	}

	token, err := h.authService.GenerateToken(userData.ID)

	if err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", http.StatusUnauthorized, "bad request", gin.H{"errors": splitError})

		c.JSON(http.StatusUnauthorized, responseError)
		return
	}
	response := helper.APIResponse("success login user", http.StatusOK, "success", gin.H{"token": token})
	c.JSON(http.StatusOK, response)
}
