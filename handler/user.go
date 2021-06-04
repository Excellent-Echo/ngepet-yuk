package handler

import (
	"net/http"
	"strconv"

	"ngepet-yuk/auth"
	"ngepet-yuk/entity"
	"ngepet-yuk/helper"
	"ngepet-yuk/layer/user"

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
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all user", 200, "status OK", users)
	c.JSON(200, response)
}

// CREATE NEW USER OR REGISTER
func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var inputUser entity.UserInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {

		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	// max len password 6
	if err := helper.ValidatePassword(inputUser.Password); err != nil {

		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"error": "error validate password length < 6"})

		c.JSON(400, responseError)
		return
	}

	newUser, err := h.userService.SaveNewUser(inputUser)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}
	response := helper.APIResponse("success create new User", 201, "Status OK", newUser)
	c.JSON(201, response)
}

// GET USER BY ID
func (h *userHandler) GetUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		responseError := helper.APIResponse("input params error", 400, "bad request", gin.H{"errors": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success get user by ID", 200, "success", user)
	c.JSON(200, response)
}

// DELETE USER BY ID
func (h *userHandler) DeleteUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	user, err := h.userService.DeleteUserByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request delete user", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success delete user by ID", 200, "success", user)
	c.JSON(200, response)
}

// UPDATE USER BY ID
func (h *userHandler) UpdateUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	var updateUserInput entity.UpdateUserInput

	if err := c.ShouldBindJSON(&updateUserInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	idParam, _ := strconv.Atoi(id)

	// ngecek id sama apa engga sama yang di inputin
	userData := int(c.MustGet("currentUser").(int))

	if idParam != userData {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user ID not authorize"})

		c.JSON(401, responseError)
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
		//splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": err.Error()})

		c.JSON(400, responseError)
		return
	}

	userData, err := h.userService.LoginUser(inputLoginUser)

	if err != nil {
		//splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 401, "bad request", gin.H{"errors": err.Error()})

		c.JSON(401, responseError)
		return
	}

	token, err := h.authService.GenerateToken(userData.ID)

	if err != nil {
		//splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 500, "bad request", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}
	response := helper.APIResponse("success login user", 200, "success", gin.H{"token": token})
	c.JSON(200, response)
}
